package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/zerolog/log"
)

type APIRequest struct {
	method      string
	url         *url.URL
	forceHttp11 bool
	debug       bool
	headers     map[string]string
	body        interface{}
	query       map[string]string
}

type APIResponse struct {
	InnerResponse *http.Response
	Data          []byte
}

func (resp *APIResponse) UnmarshalJSONTo(target interface{}) error {
	log.Debug().Str("body", resp.String()).Msg("Execute Request")
	return json.Unmarshal(resp.Data, target)
}

func (resp *APIResponse) String() string {
	return string(resp.Data)
}

type APIRequestBuilder struct {
	apiRequest APIRequest
}

func NewAPIRequestBuilder(baseUrl string, options ...APIRequestOption) (*APIRequestBuilder, error) {
	parsedURL, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	builder := &APIRequestBuilder{
		apiRequest: APIRequest{
			url:    parsedURL,
			method: "GET",
		},
	}
	for _, option := range options {
		option(&builder.apiRequest)
	}
	return builder, nil
}

type APIRequestOption func(*APIRequest)

func WithForceHttp11() APIRequestOption {
	return func(opts *APIRequest) {
		opts.forceHttp11 = true
	}
}

func WithDebug() APIRequestOption {
	return func(opts *APIRequest) {
		opts.debug = true
	}
}

func WithMethod(method string) APIRequestOption {
	return func(opts *APIRequest) {
		opts.method = method
	}
}

func WithMethodGet() APIRequestOption {
	return WithMethod("GET")
}

func WithMethodPost() APIRequestOption {
	return WithMethod("POST")
}

func WithMethodPut() APIRequestOption {
	return WithMethod("PUT")
}

func WithMethodDelete() APIRequestOption {
	return WithMethod("DELETE")
}

func WithAccept(value string) APIRequestOption {
	return WithHeader("accept", value)
}

func WithContentType(value string) APIRequestOption {
	return WithHeader("Content-Type", value)
}

func WithContentTypeJson() APIRequestOption {
	return WithContentType("application/json")
}

func WithAuthorizationBearer(token string) APIRequestOption {
	return WithHeader("Authorization", "Bearer "+token)
}

func WithHeader(key, value string) APIRequestOption {
	return func(opts *APIRequest) {
		if opts.headers == nil {
			opts.headers = make(map[string]string)
		}
		opts.headers[key] = value
	}
}

func (opts *APIRequestBuilder) WithHeader(key, value string) {
	if opts.apiRequest.headers == nil {
		opts.apiRequest.headers = make(map[string]string)
	}
	opts.apiRequest.headers[key] = value
}

func WithBody(body interface{}) APIRequestOption {
	return func(opts *APIRequest) {
		opts.body = body
	}
}

func WithPath(path string) APIRequestOption {
	return func(opts *APIRequest) {
		opts.url.Path = path
	}
}

func WithQuery(params map[string]string) APIRequestOption {
	return func(opts *APIRequest) {
		opts.query = params
	}
}

func (opts *APIRequestBuilder) WithQuery(params map[string]string) {
	opts.apiRequest.query = params
}

func (builder *APIRequestBuilder) Build() *APIRequest {
	return &builder.apiRequest
}

func (request *APIRequest) Execute() (*APIResponse, error) {
	client := &http.Client{}

	if request.forceHttp11 {
		client.Transport = &http.Transport{
			TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
		}
	}

	var req *http.Request
	var err error

	if request.body != nil {
		bodyBytes, err := json.Marshal(request.body)
		if request.debug {
			log.Debug().Bytes("body", bodyBytes).Msg("Execute Request")
		}
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(request.method, request.buildURL(), bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(request.method, request.buildURL(), nil)
		if err != nil {
			return nil, err
		}
	}

	for key, value := range request.headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &APIResponse{
		InnerResponse: resp,
		Data:          bodyBytes,
	}, nil
}

func (request *APIRequest) buildURL() string {
	query := url.Values{}
	for key, value := range request.query {
		query.Add(key, value)
	}

	url := request.url.String()
	queryString := query.Encode()
	if len(queryString) > 0 {
		if strings.ContainsRune(request.url.RawQuery, '?') {
			url = request.url.String() + "&" + queryString
		} else {
			url = request.url.String() + "?" + queryString
		}
	}

	if request.debug {
		log.Debug().Str("url", url).Msg("Execute Request")
	}
	return url
}

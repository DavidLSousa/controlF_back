package utils

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/rs/zerolog/log"
)

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback
	}
}

func GetEnvOr(key string, fallback func() string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback()
	}
}

func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func PrintError(err error) string {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		log.Error().
			Str("filename", filename).
			Int("line", line).
			Err(err).
			Msg("PrintError")
		return err.Error()
	}
	return ""
}

func PrintErrorAnd(err error, message string) string {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		log.Error().
			Str("filename", filename).
			Int("line", line).
			Err(err).
			Msg(message)
	}
	return message
}

func PrintErrorMsg(err error, message string) {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		log.Error().
			Str("filename", filename).
			Int("line", line).
			Err(err).
			Msg(message)
	}
}

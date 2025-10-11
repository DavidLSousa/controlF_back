package utils

import (
	"controlF_back/internal/domain"
	"errors"

	"github.com/go-playground/validator/v10"
)

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Should be minimum than " + fe.Param()
	case "is_url_friendly":
		return "Value is not url friendly, like: " + ParseUrlFriendly(fe.Value().(string))
	}
	return "Unknown error"
}

func GetValidationErrors(err error) []domain.ErrorDetail {
	var ve validator.ValidationErrors
	var out []domain.ErrorDetail
	if errors.As(err, &ve) {
		out = make([]domain.ErrorDetail, len(ve))
		for i, fe := range ve {
			out[i] = domain.ErrorDetail{Field: fe.Field(), Message: getErrorMsg(fe)}
		}
	}

	return out
}

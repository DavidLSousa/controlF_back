package utils

import (
	"controlF_back/internal/modules"
	"errors"
	"unicode"

	"github.com/gin-gonic/gin/binding"
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
	case "email":
		return "Invalid email format"
	case "eqfield":
		return "Fields do not match"
	case "password":
		return "Password must be at least 8 characters long and contain uppercase, lowercase, number, and special character" // Mensagem para senha forte
	}
	return "Unknown error"
}

func GetValidationErrors(err error) []modules.ErrorDetail {
	var ve validator.ValidationErrors
	var out []modules.ErrorDetail
	if errors.As(err, &ve) {
		out = make([]modules.ErrorDetail, len(ve))
		for i, fe := range ve {
			out[i] = modules.ErrorDetail{Field: fe.Field(), Message: getErrorMsg(fe)}
		}
	}

	return out
}

func validatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	if len(password) < 8 {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func SetupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validatePasswordStrength)
	}
}

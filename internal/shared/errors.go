package shared

type ErrTypes string

const (
	IsRequired   ErrTypes = "required"
	IsInvalid    ErrTypes = "invalid"
	IsExists     ErrTypes = "exists"
	NotFound     ErrTypes = "not_found"
	Unauthorized ErrTypes = "unauthorized"
	Forbidden    ErrTypes = "forbidden"
)

func ErrMessage(s string, err ErrTypes) string {
	switch err {
	case "required":
		return s + " is required"
	case "invalid":
		return s + " is invalid"
	case "exists":
		return s + " already exists"
	case "not_found":
		return s + " not found"
	case "unauthorized":
		return "unauthorized"
	case "forbidden":
		return "forbidden"
	default:
		return s + "has error"
	}
}

const (
	CodeErrorEmailAlreadyExists = "23505"
)

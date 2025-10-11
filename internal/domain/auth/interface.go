package auth

type IAuthService interface {
	login(input LoginRequest) *LoginResponse
}

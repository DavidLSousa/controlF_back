package auth

import (
	"controlF_back/internal/domain"
	"controlF_back/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService AuthService
}

func NewAuthController(service AuthService) *AuthController {
	return &AuthController{
		AuthService: service,
	}
}

// @Summary      Autentica um usuário e retorna um token JWT
// @Description  Recebe credenciais de login (email e senha) e retorna um token JWT para acesso autenticado.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials body auth.LoginRequest true "Credenciais do usuário"
// @Success      200  {object}  auth.LoginResponse
// @Failure      400  {object}  domain.ErrorResponse "Dados inválidos"
// @Failure      401  {object}  domain.ErrorResponse "Credenciais inválidas"
// @Failure      500  {object}  domain.ErrorResponse "Erro interno do servidor"
// @Router       /auth/token [post]
func (controller *AuthController) Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view := controller.AuthService.login(input)
	c.JSON(http.StatusOK, view)
}

// @Summary      Faz logout do usuário
// @Description  Invalida o token JWT do usuário, encerrando a sessão.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400  {object}  domain.ErrorResponse "Dados inválidos"
// @Failure      500  {object}  domain.ErrorResponse "Erro interno do servidor"
// @Security     BearerAuth
// @Router       /auth/logout [post]
func (controller *AuthController) Logout(c *gin.Context) {
	view := controller.AuthService.logout()
	c.JSON(http.StatusOK, view)
}

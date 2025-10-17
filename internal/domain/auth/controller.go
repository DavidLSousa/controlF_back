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

func (controller *AuthController) login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view := controller.AuthService.login(input)
	c.JSON(http.StatusOK, view)
}

func (controller *AuthController) logout(c *gin.Context) {
	view := controller.AuthService.logout()
	c.JSON(http.StatusOK, view)
}

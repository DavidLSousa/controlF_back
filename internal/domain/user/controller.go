package user

import (
	"controlF_back/internal/domain"
	"controlF_back/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	UserService UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

// @Summary      Registra um novo usuário
// @Description  Registra um novo usuário com os dados fornecidos.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user body user.UserRegister true "Dados de registro do usuário"
// @Success      201  {object}  user.UserDto
// @Failure      400  {object}  domain.ErrorResponse "Dados inválidos"
// @Failure      500  {object}  domain.ErrorResponse "Erro interno do servidor"
// @Router       /users [post]
func (controller *UserController) Register(c *gin.Context) {
	var input UserRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		out := utils.GetValidationErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{
			Error:   "Form Validation",
			Details: out,
		})
		return
	}

	view, err := controller.UserService.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusCreated, view)
}

// @Summary      Retorna um usuário
// @Description  Retorna um usuário a partir do ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        userId path string true "ID do usuário"
// @Success      200  {object}  user.UserDto
// @Failure      400  {object}  domain.ErrorResponse "ID de usuário inválido"
// @Failure      404  {object}  domain.ErrorResponse "Usuário não encontrado"
// @Failure      500  {object}  domain.ErrorResponse "Erro interno do servidor"
// @Security     BearerAuth
// @Router       /users/{userId} [get]
func (controller *UserController) Get(c *gin.Context) {
	// userId, err := uuid.Parse(c.Param("userId"))
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
	// 	return
	// }

	// view, err := controller.UserService.Get(userId)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
	// 	return
	// }

	c.JSON(http.StatusOK, "teste")
}

// @Summary      Atualiza um usuário existente
// @Description  Atualiza os dados de um usuário específico.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        userId path string true "ID do usuário"
// @Param        user body user.UserUpdate true "Dados de atualização do usuário"
// @Success      202  {object}  user.UserDto
// @Failure      400  {object}  domain.ErrorResponse "Dados inválidos ou ID de usuário inválido"
// @Failure      404  {object}  domain.ErrorResponse "Usuário não encontrado"
// @Failure      500  {object}  domain.ErrorResponse "Erro interno do servidor"
// @Security     BearerAuth
// @Router       /users/{userId} [put]
func (controller *UserController) Put(c *gin.Context) {
	var input UserUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		out := utils.GetValidationErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{
			Error:   "Form Validation",
			Details: out,
		})
		return
	}

	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view, err := controller.UserService.Update(userId, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusAccepted, view)
}

// @Summary      Atualiza a senha de um usuário
// @Description  Atualiza a senha de um usuário específico, exigindo a senha antiga e a nova senha.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        userId path string true "ID do usuário"
// @Param        password body user.UserUpdatePassword true "Dados de atualização de senha"
// @Success      202  {object}  user.UserDto
// @Failure      400  {object}  domain.ErrorResponse "Dados inválidos, ID de usuário inválido ou senha antiga incorreta"
// @Failure      404  {object}  domain.ErrorResponse "Usuário não encontrado"
// @Failure      500  {object}  domain.ErrorResponse "Erro interno do servidor"
// @Security     BearerAuth
// @Router       /users/{userId}/password [put]
func (controller *UserController) PutPassword(c *gin.Context) {
	var input UserUpdatePassword
	if err := c.ShouldBindJSON(&input); err != nil {
		out := utils.GetValidationErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{
			Error:   "Form Validation",
			Details: out,
		})
		return
	}

	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view, err := controller.UserService.UpdatePassword(userId, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusAccepted, view)
}

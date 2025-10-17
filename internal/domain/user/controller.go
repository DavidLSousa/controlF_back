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

func (controller *UserController) post(c *gin.Context) {
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

func (controller *UserController) get(c *gin.Context) {
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

func (controller *UserController) put(c *gin.Context) {
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

func (controller *UserController) putPassword(c *gin.Context) {
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

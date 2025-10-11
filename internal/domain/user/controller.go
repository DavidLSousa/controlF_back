package user

import (
	"controlF_back/internal/domain"
	"controlF_back/internal/models"
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

func (controller *UserController) list(c *gin.Context) {
	userIdStr := c.Query("userId") // pega userId da query, se necessário
	if userIdStr == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "userId is required"})
		return
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "invalid userId"})
		return
	}

	p := models.NewPagination(c)
	result, err := controller.UserService.List(userId, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
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
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view, err := controller.UserService.Get(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusOK, view)
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

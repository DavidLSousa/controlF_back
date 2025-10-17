package user

import (
	"controlF_back/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, controller UserController) {
	unprotected := r.Group("/api")
	unprotected.POST("/users", controller.Register)

	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("/users/:userId", controller.Get)
	protected.PUT("/users/:userId", controller.Put)
	protected.PUT("/users/:userId/password", controller.PutPassword)
}

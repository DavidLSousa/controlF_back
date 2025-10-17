package user

import (
	"controlF_back/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, controller UserController) {
	unprotected := r.Group("/api")
	unprotected.POST("/users", controller.register)

	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("/users/:userId", controller.get)
	protected.PUT("/users/:userId", controller.put)
	protected.PUT("/users/:userId/password", controller.putPassword)
}

package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, controller AuthController) {
	unprotected := r.Group("/api")
	unprotected.POST("/auth/token", controller.login)

	protected := r.Group("/api")
	// protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/auth/logout", controller.logout)
}

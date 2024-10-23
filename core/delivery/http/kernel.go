package http

import (
	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/core/delivery/handlers"
	"github.com/haisabdillah/golang-auth/core/delivery/middleware"
)

// SetupRoutes sets up the HTTP routes for the application.
func SetupRoutes(r *gin.Engine, handler *handlers.Handler) {
	guestRouter(r, handler)
	authGroup := r.Group("/", middleware.AuthMiddleware())
	authRouter(authGroup, handler)
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/core/delivery/handlers"
)

func authRouter(r *gin.RouterGroup, h *handlers.Handler) {
	r.POST("/users", h.UserCreate)
	r.GET("/auth/me", h.AuthMe)
	r.PUT("/auth/profile", h.AuthProfile)
	r.GET("/auth/refresh-token", h.AuthRefreshToken)
	r.POST("/auth/change-password", h.AuthChangePassword)
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/core/delivery/handlers"
)

func guestRouter(r *gin.Engine, h *handlers.Handler) {
	r.POST("/roles", h.RoleCreate)
	r.PUT("/roles/:id", h.RoleUpdate)
	r.DELETE("/roles/:id", h.RoleDelete)
	r.GET("/roles", h.RoleGet)
	// Define authentication routes
	r.POST("/auth/login", h.AuthLogin)
	r.POST("/permissions", h.PermissionCreate)
	r.GET("/permissions", h.PermissionGet)
	r.GET("/permissions/:id", h.PermissionFirst)
	r.PUT("/permissions/:id", h.PermissionUpdate)
	r.DELETE("/permissions/:id", h.PermissionDelete)

}

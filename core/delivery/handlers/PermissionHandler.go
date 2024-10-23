package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/core/dto"
)

func (h *Handler) PermissionGet(c *gin.Context) {
	result, err := h.service.PermissionGet()
	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", result)
}

func (h *Handler) PermissionFirst(c *gin.Context) {
	id := c.Param("id")
	result, err := h.service.PermissionFirst(id)
	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", result)

}

func (h *Handler) PermissionCreate(c *gin.Context) {
	body := dto.PermissionDto{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ResponseInvalidBindJson(c, err)
		return
	}
	if err := h.service.PermissionCreate(&body); err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", nil)

}
func (h *Handler) PermissionUpdate(c *gin.Context) {
	id := c.Param("id")
	body := dto.PermissionDto{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ResponseInvalidBindJson(c, err)
		return
	}
	if err := h.service.PermissionUpdate(id, &body); err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", nil)

}
func (h *Handler) PermissionDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.PermissionDelete(id); err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", nil)

}

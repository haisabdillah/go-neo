package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/core/dto"
)

func (h *Handler) RoleCreate(c *gin.Context) {
	body := dto.RoleDto{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ResponseInvalidBindJson(c, err)
		return
	}
	if err := h.service.RoleCreate(&body); err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", nil)

}

func (h *Handler) RoleUpdate(c *gin.Context) {
	id := c.Param("id")
	body := dto.RoleDto{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ResponseInvalidBindJson(c, err)
		return
	}
	if err := h.service.RoleUpdate(id, &body); err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", nil)

}

func (h *Handler) RoleDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.RoleDelete(id); err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", nil)
}

func (h *Handler) RoleGet(c *gin.Context) {
	data, err := h.service.RoleGet()
	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", data)

}

func (h *Handler) RoleFirst(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.RoleFirst(id)
	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", data)

}

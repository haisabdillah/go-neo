package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/core/dto"
)

func (h *Handler) UserCreate(c *gin.Context) {
	body := dto.UserDto{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ResponseInvalidBindJson(c, err)
		return
	}
	if err := h.service.UserCreate(&body); err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", nil)
}

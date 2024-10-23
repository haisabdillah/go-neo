package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/core/dto"
)

func (h *Handler) AuthLogin(c *gin.Context) {
	body := dto.AuthLoginDto{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ResponseInvalidBindJson(c, err)
		return
	}
	token, err := h.service.AuthLogin(&body)
	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", token)
}

func (h *Handler) AuthMe(c *gin.Context) {
	authID := c.GetUint("authID")
	data, err := h.service.AuthMe(authID)
	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", data)
}

func (h *Handler) AuthRefreshToken(c *gin.Context) {
	authID := c.GetUint("authID")
	data, err := h.service.AuthRefreshToken(authID)
	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", data)
}

func (h *Handler) AuthProfile(c *gin.Context) {
	authID := c.GetUint("authID")
	body := dto.AuthProfileDto{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ResponseInvalidBindJson(c, err)
		return
	}
	err := h.service.AuthProfile(authID, &body)
	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", nil)
}

func (h *Handler) AuthChangePassword(c *gin.Context) {
	authID := c.GetUint("authID")

	body := dto.AuthChangePasswordDto{}
	if err := c.ShouldBindJSON(&body); err != nil {
		ResponseInvalidBindJson(c, err)
		return
	}

	if err := h.service.AuthChangePassword(authID, &body); err != nil {
		ResponseError(c, err)
		return
	}
	ResponseOK(c, "OK", nil)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/core/services"
	myError "github.com/haisabdillah/golang-auth/pkg/errors"
)

type Handler struct {
	service *services.Service
}

func NewHandler(serv *services.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

func ResponseError(c *gin.Context, err error) {
	type responseError struct {
		Message string `json:"message"`
		Error   any    `json:"error,omitempty"`
		Errors  any    `json:"errors,omitempty"`
	}
	if customError, ok := err.(myError.Err); ok {
		resp := responseError{
			Message: customError.Message,
			Errors:  customError.Errors,
		}
		if customError.StatusCode == 500 {
			jsonData, err := json.Marshal(customError.Errors)
			if err != nil {
				log.Fatalf("Error marshaling to JSON: %v", err)
			}
			c.Set("error", string(jsonData))
			fmt.Println(string(jsonData))
			resp = responseError{
				Message: customError.Error(),
			}
		}
		c.JSON(customError.StatusCode, resp)
	} else {
		resp := responseError{
			Message: "Undefined Error",
			Error:   err.Error(),
		}
		c.JSON(500, resp)
	}
}

func ResponseInvalidBindJson(c *gin.Context, err error) {
	type responseError struct {
		Message string `json:"message"`
		Error   any    `json:"error,omitempty"`
	}
	resp := responseError{
		Message: "Invalid bind JSON",
		Error:   err.Error(),
	}
	c.JSON(400, resp)
}

func ResponseOK(c *gin.Context, message string, data interface{}) {
	type responseOK struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
	responseMessage := "OK"
	if message != "" {
		responseMessage = message
	}
	resp := responseOK{
		Message: responseMessage,
		Data:    data,
	}
	c.JSON(200, resp)
}

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusOKWithData struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type StatusOK struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type StatusError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type StatusErrorCustom struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func ResponseOKWithData(c *gin.Context, data interface{}) {
	response := StatusOKWithData{
		Code:    1000,
		Data:    data,
		Message: "OK",
	}

	c.JSON(http.StatusOK, response)
}

func ResponseCreated(c *gin.Context, data interface{}) {
	response := StatusOKWithData{
		Code:    1000,
		Data:    data,
		Message: "Created",
	}

	c.JSON(http.StatusCreated, response)
}

func ResponseOK(c *gin.Context, message string) {
	response := StatusOK{
		Code:    1000,
		Message: message,
	}

	c.JSON(http.StatusOK, response)
}

func ResponseError(c *gin.Context, err string, code int) {
	response := StatusError{
		Code:    99,
		Message: err,
	}

	c.JSON(code, response)
}

func ResponseCustomError(c *gin.Context, err interface{}, code int) {
	response := StatusErrorCustom{
		Code:    99,
		Message: err,
	}

	c.JSON(code, response)
}

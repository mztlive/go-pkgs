package httpresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mztlive/project-template/pkg/response"
)

func UnAuthorized(c *gin.Context, message string) {
	Fail(c, http.StatusUnauthorized, http.StatusUnauthorized, message)
}

func UnPermission(c *gin.Context, message string) {
	Fail(c, http.StatusForbidden, http.StatusForbidden, message)
}

func BadRequest(c *gin.Context, message string) {
	Fail(c, http.StatusBadRequest, http.StatusBadRequest, message)
}

func SystemError(c *gin.Context, message string) {
	Fail(c, http.StatusInternalServerError, http.StatusInternalServerError, message)
}

func Fail(c *gin.Context, httpCode int, code int, message string) {
	apiResponse := response.ApiResponse{
		HttpCode: httpCode,
		Code:     code,
		Msg:      message,
		Data: struct {
		}{},
	}
	makeResponseJson(c, &apiResponse)
}

// Success 成功的请求
func Success(c *gin.Context, data interface{}) {
	var apiResponse response.ApiResponse
	if data != nil {
		apiResponse = response.ApiResponse{
			Code: http.StatusOK,
			Msg:  "success",
			Data: data,
		}
	} else {
		apiResponse = response.ApiResponse{
			Code: http.StatusOK,
			Msg:  "success",
			Data: struct {
			}{},
		}
	}
	makeResponseJson(c, &apiResponse)
}

func makeResponseJson(c *gin.Context, apiResponse *response.ApiResponse) {
	resp := gin.H{
		"meta": gin.H{
			"code": apiResponse.Code,
			"msg":  apiResponse.Msg,
		},
		"data": apiResponse.Data,
	}
	responseJson(c, apiResponse.HttpCode, resp)
	return
}

func responseJson(c *gin.Context, httpCode int, resp interface{}) {
	c.JSON(httpCode, resp)
	c.Abort()
	return
}

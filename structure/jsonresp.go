package structure

import "net/http"

// JsonApiResponse 基础的HttpApi返回结构
type JsonApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"result"`
}

// ForbiddenRequest 没有权限
func ForbiddenRequest() JsonApiResponse {
	return JsonApiResponse{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
		Data:    nil,
	}
}

// BadRequest 错误的请求, 一般是逻辑上的，比如说传少参数了
func BadRequest(message error) JsonApiResponse {
	response := JsonApiResponse{
		Code:    http.StatusBadRequest,
		Message: message.Error(),
		Data:    nil,
	}

	return response
}

// Success 成功的请求
func Success(data interface{}) JsonApiResponse {
	return JsonApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

// SystemError 系统错误
func SystemError(message error) JsonApiResponse {
	return JsonApiResponse{
		Code:    http.StatusInternalServerError,
		Message: message.Error(),
		Data:    nil,
	}
}

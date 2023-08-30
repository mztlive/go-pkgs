package response

// ApiResponse 基础的HttpApi返回结构
type ApiResponse struct {
	HttpCode int         `json:"httpCode"`
	Code     int         `json:"code"`
	Msg      string      `json:"Msg"`
	Data     interface{} `json:"data"`
}

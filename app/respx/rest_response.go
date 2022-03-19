package respx

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Reason  string      `json:"reason,omitempty"`
}

func NewSuccessEmptyResponse() *Response {
	return &Response{
		Code: 200,
		Message: "success",
	}
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Code: 200,
		Message: "success",
		Data: data,
	}
}
package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error"`
	Data    any    `json:"data"`
}

func BuildResponseSuccess(status bool, message string, data any) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

func BuildResponseFailed(message string, err string, data any) Response {
	res := Response{
		Status: false,
		Message: message,
		Error: nil,
		Data: data,
	}
	return res
}
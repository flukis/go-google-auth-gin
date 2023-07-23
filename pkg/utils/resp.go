package utils

type ApiResponse struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

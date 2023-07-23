package utils

type ErrorResponse struct {
	Code  int   `json:"code"`
	Error error `json:"error"`
}

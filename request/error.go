package request

import "net/http"

type RequestError struct {
	StatusCode int

	Error error
}
/*
func NewError(StatusCode int, err error) RequestError {
	return RequestError{StatusCode: StatusCode, Error: err}
}*/
func NewRequestError(err error) RequestError {
	return RequestError{StatusCode: http.StatusBadRequest, Error: err}
}
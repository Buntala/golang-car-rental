package request

import "net/http"

type RequestError struct {
	StatusCode int

	E error
}

func (reqErr RequestError) Error() string{
	return reqErr.E.Error()
}

func NewRequestError(err error) RequestError {
	return RequestError{StatusCode: http.StatusBadRequest, E: err}
}
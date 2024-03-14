package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse[_] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func InternalServerError(msg string) Response[any] {
	return errors[any](msg, http.StatusInternalServerError)
}

func NotFound(msg string) Response[any] {
	return errors[any](msg, http.StatusNotFound)
}

func Unauthorized(msg string) Response[any] {
	return errors[any](msg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response[any] {
	return errors[any](msg, http.StatusForbidden)
}

func BadRequest(msg string) Response[any] {
	return errors[any](msg, http.StatusBadRequest)
}

func errors[_](msg string, status int) Response[any] {
	return &ErrorResponse{
		Status:  status,
		Message: msg,
	}
}

func (e ErrorResponse[_]) Error() string {
	return e.Message
}

func (e ErrorResponse[_]) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse[_]) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ErrorResponse[_]) GetData() interface{} {
	return nil
}

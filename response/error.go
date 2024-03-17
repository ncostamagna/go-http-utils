package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func InternalServerError(msg string) Response[any, ErrorResponse] {
	return errors(msg, http.StatusInternalServerError)
}

func NotFound(msg string) Response[any, ErrorResponse] {
	return errors(msg, http.StatusNotFound)
}

func Unauthorized(msg string) Response[any, ErrorResponse] {
	return errors(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response[any, ErrorResponse] {
	return errors(msg, http.StatusForbidden)
}

func BadRequest(msg string) Response[any, ErrorResponse] {
	return errors(msg, http.StatusBadRequest)
}

func errors(msg string, status int) Response[any, ErrorResponse] {
	return Response[any, ErrorResponse]{
		Body: ErrorResponse{
			Status:  status,
			Message: msg,
		},
	}
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

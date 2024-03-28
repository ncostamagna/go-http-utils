package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	badRequestMsg          = "Your request is in a bad format."
	forbiddenMsg           = "You are not authorized to perform the requested action."
	unauthorizedMsg        = "You are not authenticated to perform the requested action."
	notFoundMsg            = "The requested resource was not found."
	internalServerErrorMsg = "We encountered an error while processing your request."
	invalidInputMsg        = "There is some problem with the data you submitted."
)

func InternalServerError(msg string) Response {
	return errors(msg, internalServerErrorMsg, http.StatusInternalServerError)
}

func NotFound(msg string) Response {
	return errors(msg, notFoundMsg, http.StatusNotFound)
}

func Unauthorized(msg string) Response {
	return errors(msg, unauthorizedMsg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response {
	return errors(msg, forbiddenMsg, http.StatusForbidden)
}

func BadRequest(msg string) Response {
	return errors(msg, badRequestMsg, http.StatusBadRequest)
}

func errors(msg string, defMsg string, status int) Response {
	if msg == "" {
		msg = defMsg
	}

	return &ErrorResponse{
		Status:  status,
		Message: msg,
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

func (e *ErrorResponse) GetData() interface{} {
	return nil
}

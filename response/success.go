package response

import (
	"encoding/json"
	"github.com/ncostamagna/go-http-utils/meta"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Meta    *meta.Meta  `json:"meta,omitempty"`
}

const (
	successMsg              = "Ok request."
	createdMsg              = "Created request."
	acceptedMsg             = "Accepted request."
	nonAuthoritativeInfoMsg = "Non-Authoritative Information request."
	noContentMsg            = "No Content request."
	resetContentMsg         = "Reset Content request."
	partialContentMsg       = "Partial Content request."
)

func OK(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, successMsg, data, meta, http.StatusOK)
}

func Created(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, createdMsg, data, meta, http.StatusCreated)
}

func Accepted(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, acceptedMsg, data, meta, http.StatusAccepted)
}

func NonAuthoritativeInfo(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, nonAuthoritativeInfoMsg, data, meta, http.StatusNonAuthoritativeInfo)
}

func NoContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, noContentMsg, data, meta, http.StatusNoContent)
}

func ResetContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, resetContentMsg, data, meta, http.StatusResetContent)
}

func PartialContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, partialContentMsg, data, meta, http.StatusPartialContent)
}

func success(msg, defMsg string, data interface{}, meta *meta.Meta, code int) Response {

	if msg == "" {
		msg = defMsg
	}
	return &SuccessResponse{
		Message: msg,
		Status:  code,
		Data:    &data,
		Meta:    meta,
	}
}

func (s *SuccessResponse) Error() string {
	return ""
}

func (s *SuccessResponse) StatusCode() int {
	return s.Status
}

func (s *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}

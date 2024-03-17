package response

import (
	"encoding/json"
	"github.com/ncostamagna/go_lib_response/meta"
	"net/http"
)

type SuccessResponse[D any] struct {
	Message string     `json:"message"`
	Status  int        `json:"status"`
	Data    *D         `json:"data"`
	Meta    *meta.Meta `json:"meta,omitempty"`
}

func OK[D any](msg string, data D, meta *meta.Meta) *Response[D, SuccessResponse[D]] {
	return success(msg, data, meta, http.StatusOK)
}

func Created[D any](msg string, data D, meta *meta.Meta) *Response[D, SuccessResponse[D]] {
	return success(msg, data, meta, http.StatusCreated)
}

func Accepted[D any](msg string, data D, meta *meta.Meta) *Response[D, SuccessResponse[D]] {
	return success(msg, data, meta, http.StatusAccepted)
}

func NonAuthoritativeInfo[D any](msg string, data D, meta *meta.Meta) *Response[D, SuccessResponse[D]] {
	return success(msg, data, meta, http.StatusNonAuthoritativeInfo)
}

func NoContent[D any](msg string, data D, meta *meta.Meta) *Response[D, SuccessResponse[D]] {
	return success(msg, data, meta, http.StatusNoContent)
}

func ResetContent[D any](msg string, data D, meta *meta.Meta) *Response[D, SuccessResponse[D]] {
	return success(msg, data, meta, http.StatusResetContent)
}

func PartialContent[D any](msg string, data D, meta *meta.Meta) *Response[D, SuccessResponse[D]] {
	return success(msg, data, meta, http.StatusPartialContent)
}

func success[D any](msg string, data D, meta *meta.Meta, code int) *Response[D, SuccessResponse[D]] {
	return &Response[D, SuccessResponse[D]]{
		Body: SuccessResponse[D]{
			Message: msg,
			Status:  code,
			Data:    &data,
			Meta:    meta,
		},
	}
}

func (s *SuccessResponse[_]) Error() string {
	return ""
}

func (s *SuccessResponse[_]) StatusCode() int {
	return s.Status
}

func (s *SuccessResponse[_]) GetBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SuccessResponse[D]) GetData() *D {
	return s.Data
}

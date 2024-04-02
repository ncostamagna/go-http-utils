package request

import (
	"errors"
)

var ErrFormatValue = errors.New("error with the format value in decode request")
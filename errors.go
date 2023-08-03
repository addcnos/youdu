package youdu

import "fmt"

type Error struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func newError(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

package errors

import "fmt"

type httpError struct {
	code int
	msg  string
}

func (e *httpError) Error() string {
	return fmt.Sprintf("[ %d ] %s", e.code, e.msg)
}

func newHttpError(code int, msg string) *httpError {
	return &httpError{code: code, msg: msg}
}

func ParseError(err error) (int, string) {
	if err == nil {
		return OK.code, OK.msg
	}
	if e, ok := err.(*httpError); ok {
		return e.code, e.msg
	}
	return InternalError.code, InternalError.msg
}

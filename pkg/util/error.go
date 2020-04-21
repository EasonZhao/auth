package util

import "github.com/securityin/auth/pkg/e"

// Error 带有错误码 与 错误信息的错误类
type Error interface {
	error
	Code() int
}

// ErrNew returns an error that formats as the given text.
func ErrNew(code int, text string) Error {
	return &errorString{code, text}
}

// ErrNewCode returns an error that formats as the given text.
func ErrNewCode(code int) Error {
	return &errorString{code, e.GetMsg(code)}
}

// ErrNewSQL returns an error that formats as the given text.
func ErrNewSQL(err error) Error {
	return &errorString{e.ErrorExecSql, err.Error()}
}

// ErrNewErr returns an error that formats as the given text.
func ErrNewErr(err error) Error {
	return &errorString{e.ERROR, err.Error()}
}

// errorString is a trivial implementation of error.
type errorString struct {
	code int
	s    string
}

func (e *errorString) Error() string {
	return e.s
}

func (e *errorString) Code() int {
	return e.code
}

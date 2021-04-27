package e

import "github.com/pkg/errors"

type WithCoder interface {
	Code(code int, extras ...error) error                          // set code
	Msg(code int, s string, extras ...error) error                 // set code and msg
	Err(code int, err error, extras ...error) error              // set code and err
	ErrMsg(code int, err error, s string, extras ...error) error // set code, err and msg
}
type withCode struct {
}

// Invoke this function
func WithCode() WithCoder {
	return &withCode{}
}
func (w *withCode) Code(code int, extras ...error) error {
	return &StackError{
		code:      code,
		err: errors.New(""),
		extraErrs: extras,
	}
}

func (w *withCode) Msg(code int, s string, extras ...error) error {
	return &StackError{
		code:      code,
		err: errors.New(""),
		msg:       s,
		extraErrs: extras,
	}
}

func (w *withCode) Err(code int, err error, extras ...error) error {
	return &StackError{
		code:      code,
		err:       errors.WithStack(err),
		extraErrs: extras,
	}
}

func (w *withCode) ErrMsg(code int, err error, s string, extras ...error) error {
	return &StackError{
		code:      code,
		msg:       s,
		err: errors.WithStack(err),
		extraErrs: extras,
	}
}

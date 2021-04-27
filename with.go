package e

import "github.com/pkg/errors"

type Wither interface {
	Msg(s string, extras ...error) error               // set msg
	Err(err error, extras ...error) error              // set error
	ErrMsg(err error, s string, extras ...error) error // set error and msg
}
type with struct {
}

// Invoke this function
func With() Wither {
	return &with{}
}

// Set the error when setting msg. because heap tracing is implement.
func (w *with) Msg(s string, extras ...error) error {
	return &StackError{
		err:       errors.New(""),
		msg:       s,
		extraErrs: extras,
	}
}

func (w *with) Err(err error, extras ...error) error {
	return &StackError{
		err:       errors.WithStack(err),
		extraErrs: extras,
	}
}

func (w *with) ErrMsg(err error, s string, extras ...error) error {
	return &StackError{
		err:       errors.WithStack(err),
		msg:       s,
		extraErrs: extras,
	}
}

package e

import "github.com/pkg/errors"

// Code and uuid can be set.

type WithUUIDer interface {
	UUID(uuid string, extras ...error) error                                  // set uuid
	Code(code int, uuid string, extras ...error) error                        // set code
	Msg(code int, uuid string, s string, extras ...error) error               // set code and msg
	Err(code int, uuid string, err error, extras ...error) error              // set code and err
	ErrMsg(code int, uuid string, err error, s string, extras ...error) error // set code, err and msg
}
type withUUID struct {
}

// Invoke this function
func WithUUID() WithUUIDer {
	return &withUUID{}
}
func (w *withUUID) UUID(uuid string, extras ...error) error {
	return &StackError{
		uuid:      uuid,
		err:       errors.New(""),
		extraErrs: extras,
	}
}
func (w *withUUID) Code(code int, uuid string, extras ...error) error {
	return &StackError{
		code:      code,
		uuid:      uuid,
		err:       errors.New(""),
		extraErrs: extras,
	}
}

func (w *withUUID) Msg(code int, uuid string, s string, extras ...error) error {
	return &StackError{
		code:      code,
		uuid:      uuid,
		msg:       s,
		err:       errors.New(""),
		extraErrs: extras,
	}
}

func (w *withUUID) Err(code int, uuid string, err error, extras ...error) error {
	return &StackError{
		code:      code,
		uuid:      uuid,
		err:       errors.WithStack(err),
		extraErrs: extras,
	}
}

func (w *withUUID) ErrMsg(code int, uuid string, err error, s string, extras ...error) error {
	return &StackError{
		code:      code,
		uuid:      uuid,
		msg:       s,
		err:       errors.WithStack(err),
		extraErrs: extras,
	}
}

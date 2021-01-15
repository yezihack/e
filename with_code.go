package e

import "github.com/pkg/errors"

// 信息+额外errs
func NewCode(code int, s string, extras ...error) error {
	return &StackError{
		code:      code,
		err:       errors.New(s),
		extraErrs: extras,
	}
}

// 格式+信息
func ErrorCodeF(code int, format string, args ...interface{}) error {
	return &StackError{
		code: code,
		err:  errors.Errorf(format, args...),
	}
}

// 错误堆栈设置+额外errs
func WithCodeStack(code int, err error, extras ...error) error {
	return &StackError{
		code:      code,
		err:       errors.WithStack(err),
		extraErrs: extras,
	}
}

// 错误+信息+额外errs
func WithCodeMessage(code int, err error, msg string, extras ...error) error {
	err = errors.WithStack(err)
	return &StackError{
		code:      code,
		err:       errors.Wrap(err, msg),
		extraErrs: extras,
	}
}

// 错误+格式+信息
func WithCodeMessageF(code int, err error, format string, args ...interface{}) error {
	return &StackError{
		code: code,
		err:  errors.Wrapf(err, format, args...),
	}
}

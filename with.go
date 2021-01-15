package e

import "github.com/pkg/errors"

// 信息+额外errs
func New(s string, extras ...error) error {
	return &StackError{
		code:      0,
		err:       errors.New(s),
		extraErrs: extras,
	}
}

// 格式+信息
func ErrorF(format string, args ...interface{}) error {
	return &StackError{
		code: 0,
		err:  errors.Errorf(format, args...),
	}
}

// 错误堆栈设置+额外errs
func WithStack(err error, extras ...error) error {
	return &StackError{
		code:      0,
		err:       errors.WithStack(err),
		extraErrs: extras,
	}
}

// 错误+信息+额外errs
func WithMessage(err error, msg string, extras ...error) error {
	err = errors.WithStack(err)
	return &StackError{
		code:      0,
		err:       errors.Wrap(err, msg),
		extraErrs: extras,
	}
}

// 错误+格式+信息
func WithMessageF(err error, format string, args ...interface{}) error {
	return &StackError{
		code: 0,
		err:  errors.Wrapf(err, format, args...),
	}
}

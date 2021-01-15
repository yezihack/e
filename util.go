package e

import (
	"fmt"

	"github.com/pkg/errors"
)

// 对 error 进行断言, 是否自定义的错误,即StackError
func Assert(e error) *StackError {
	switch e.(type) {
	case *StackError:
		return e.(*StackError)
	default:
		return nil
	}
}

// 对 error 进行断言, 是否自定义的错误,即StackError
// 判断是否是,返回二个参数, 带bool
func AssertB(e error) (*StackError, bool) {
	switch e.(type) {
	case *StackError:
		return e.(*StackError), true
	default:
		return nil, false
	}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
func As(err, target error) bool {
	return errors.As(err, target)
}

func UnWrap(err error) error {
	return errors.Unwrap(err)
}

// 打印堆栈信息 by str
func ToStr(err error) string {
	stacks := MarshalStack(err)
	if stacks == nil {
		return err.Error()
	}
	var s string
	for idx, item := range stacks {
		if idx == 0 {
			continue
		}
		s += fmt.Sprintf("file:%s, line:%s, func:%s\n", item.File, item.LineCode, item.FuncName)
	}
	return s
}

// 打印堆栈信息 by slice
func ToArr(err error) []string {
	stacks := MarshalStack(err)
	if stacks == nil {
		return []string{err.Error()}
	}
	result := make([]string, len(stacks))
	for idx, item := range stacks {
		result[idx] = fmt.Sprintf("file:%s, line:%s, func:%s", item.File, item.LineCode, item.FuncName)
	}
	return result
}

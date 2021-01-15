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
	for _, item := range stacks {
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

package e

import (
	"fmt"
)

// 转换成自定义 error 对象
func Convert(e error) *StackError {
	switch e.(type) {
	case *StackError:
		return e.(*StackError)
	default:
		return nil
	}
}

// 对 error 进行断言, 是否自定义的错误
func Assert(e error) bool {
	switch e.(type) {
	case *StackError:
		return true
	default:
		return false
	}
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

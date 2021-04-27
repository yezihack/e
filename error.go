package e

import "strings"

// 错误自定义, 支持堆栈信息的错误结构体
type StackError struct {
	code      int     // 错误码
	err       error   // 错误信息 by error
	msg       string  // 消息
	uuid      string  // 跟踪ID
	extraErrs []error // 额外的 error
}
// 获取错误码. 只有设置了才会有.
func (c *StackError) Code() int {
	return c.code
}

// 继承系统定义的 error 接口
func (c *StackError) Error() string {
	return c.err.Error()
}

// return uuid
func (c *StackError) UUID() string {
	return c.uuid
}

// 获取错误信息
func (c *StackError) Msg() string {
	return c.msg
}

// 获取错误
func (c *StackError) Err() error {
	return c.err
}

// 获取扩展的 error 列表
func (c *StackError) Errs() []error {
	return c.extraErrs
}

// 判断是否存在扩展错误信息
func (c *StackError) ExistExtra() bool {
	if len(c.extraErrs) == 0 {
		return false
	}
	return true
}

// 获取扩展的 error 拼接的字符串
func (c *StackError) ToStrByExtra() string {
	s := ""
	for _, err := range c.extraErrs {
		s += err.Error() + ","
	}
	s = strings.TrimRight(s, ",")
	return s
}

// 打印堆栈信息 by string, 带 \n 分隔
func (c *StackError) ToStr() string {
	return ToStr(c.err)
}

// 打印堆栈信息 by array
func (c *StackError) ToArr() []string {
	return ToArr(c.err)
}

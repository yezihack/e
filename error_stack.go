package e

import "github.com/pkg/errors"

// 定义输出的结构体
type StackSourceEntity struct {
	File     string `json:"file"` // 文件名称
	LineCode string `json:"line"` // 行号
	FuncName string `json:"func"` // 函数名称
}

// 状态结构体
type state struct {
	b []byte
}

// Write implement fmt.Formatter interface.
func (s *state) Write(b []byte) (n int, err error) {
	s.b = b
	return len(b), nil
}

// Width implement fmt.Formatter interface.
func (s *state) Width() (wid int, ok bool) {
	return 0, false
}

// Precision implement fmt.Formatter interface.
func (s *state) Precision() (prec int, ok bool) {
	return 0, false
}

// Flag implement fmt.Formatter interface.
func (s *state) Flag(c int) bool {
	return false // 控制是否输出详细路径. 如果为false则输出相对路径, 如果是true输出绝对路径
}

func frameField(f errors.Frame, s *state, c rune) string {
	f.Format(s, c)

	return string(s.b)
}

// MarshalStack implements pkg/errors stack trace marshaling.
func MarshalStack(err error) []*StackSourceEntity {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	sterr, ok := err.(stackTracer)
	if !ok {
		return nil
	}
	st := sterr.StackTrace()
	s := &state{}
	out := make([]*StackSourceEntity, 0, len(st))
	size := len(st)
	for idx, frame := range st {
		if idx == 0 || size-idx <= 2 {
			continue
		}
		out = append(out, &StackSourceEntity{
			File:     frameField(frame, s, 's'),
			LineCode: frameField(frame, s, 'd'),
			FuncName: frameField(frame, s, 'n'),
		})
	}
	return out
}

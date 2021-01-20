package main

import (
	"fmt"
	"testing"

	"github.com/yezihack/e"
)

// 如果使用带 code 的 error
func fooCode() error {
	return e.NewCode(400, "eoo error")
}
func TestFooCode(t *testing.T) {
	err := fooCode()
	if e.Assert(err) {
		e1 := e.Convert(err)
		fmt.Printf("code:%d, err:%s\n", e1.Code(), e1.Msg())
	}
	//out:
	//code:400, err:eoo error
	if e.Assert(err) {
		e1 := e.Convert(err)
		fmt.Printf("code:%d, err:%s\n", e1.Code(), e1.ToStr())
	}
	// out:
	//code:400, err:file:2.code.go, line:11, func:fooCode
	//file:2.code.go, line:15, func:main
}

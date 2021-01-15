package main

import (
	"fmt"
	"log"

	"github.com/yezihack/e"
)

func foo(s string) error {
	return e.New("foo," + s)
}
func GetFoo() error {
	return foo("my is foo")
}

func main() {
	// 普通使用
	err := GetFoo()
	if err != nil {
		log.Println(err)
	}
	// out:
	// 2021/01/15 19:59:17 foo,my is foo
	// 输出堆栈信息
	e1 := e.Assert(err)
	if e1 != nil {
		fmt.Println(e1.ToStr())
	}
	// out:
	// file:with.go, line:9, func:New
	// file:1.how.go, line:11, func:foo
	// file:1.how.go, line:14, func:GetFoo
	// file:1.how.go, line:19, func:main
	// file:proc.go, line:204, func:main
	// file:asm_amd64.s, line:1374, func:goexit
}

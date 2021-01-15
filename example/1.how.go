package main

import (
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
	// (1)普通使用
	err := GetFoo()
	if err != nil {
		log.Println(err)
	}
	// out:
	// 2021/01/15 19:59:17 foo,my is foo

	// (2)输出堆栈信息 by string
	errStr := e.Assert(err) // 需要判断是否是自定义error, 否则无法输出堆栈信息.
	if errStr != nil {
		log.Println(errStr.ToStr())
	}
	// out:
	//2021/01/15 20:15:03 file:1.how.go, line:10, func:foo
	//file:1.how.go, line:13, func:GetFoo
	//file:1.how.go, line:18, func:main

	// (3)输出堆栈信息 by array
	errArr := e.Assert(err) // 需要判断是否是自定义error, 否则无法输出堆栈信息.
	if errArr != nil {
		log.Println(errArr.ToArr())
	}
	// out
	// 2021/01/15 20:20:05 [file:with.go, line:9, func:New file:1.how.go, line:10, func:foo file:1.how.go, line:13, func:GetFoo file:1.how.go, line:18, func:main file:proc.go, line:204, func:main file:asm_amd64.s, line:1374, func:goexit]
}

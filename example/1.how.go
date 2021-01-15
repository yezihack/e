package main

import (
	"log"

	"github.com/yezihack/e"
)

func foo(s string) error {
	return e.New("foo," + s)
}

func main() {
	// (1)普通使用
	err := foo("stack error")
	if err != nil {
		log.Println(err)
	}
	// out:
	// 2021/01/15 20:23:21 foo,stack error

	// (2)输出堆栈信息 by string
	errStr := e.Assert(err) // 需要判断是否是自定义error, 否则无法输出堆栈信息.
	if errStr != nil {
		log.Println(errStr.ToStr())
	}
	// out:
	//2021/01/15 20:23:21 file:1.how.go, line:10, func:foo
	//file:1.how.go, line:15, func:main

	// (3)输出堆栈信息 by array
	errArr := e.Assert(err) // 需要判断是否是自定义error, 否则无法输出堆栈信息.
	if errArr != nil {
		log.Println(errArr.ToArr())
	}
	// out
	//2021/01/15 20:23:21 [file:1.how.go, line:10, func:foo file:1.how.go, line:15, func:main]
}

# golang 追踪堆栈错误包
> Golang tracks stack error package. 追踪堆栈错误包

## 介绍(Introduction)
基于`github.com/pkg/errors`包进行封装,方便开发项目中使用. 

这个包本身是可以输出堆栈信息的,但是输出是原始的堆栈信息

存储在日志里查看不友好,而`github.com/yezihack/e`对堆栈信息进行封装了.

而且重新实现了系统自带的`error`接口,提供 `code` 自定义功能.

## 文档(Documentation)
[https://godoc.org/github.com/yezihack/e](https://godoc.org/github.com/yezihack/e)

## 安装(install)
`go get github.com/yezihack/e`

## 简单使用
```go
func foo() error {
	return e.New("foo")
}
func main() {
    err := foo()
    if err != nil && e.Assert(err) { // 需要判断是否是自定义error, 否则无法输出堆栈信息.
        log.Println(e.Convert(err).ToStr()) // 输出字符串形式
        log.Println(e.Convert(err).ToArr()) // 输出数组形式
    }
}
```

## 实例(Example)
1. [基本用法](example/1.how.go)
1. [Code用法](example/2.code.go)
1. [兼容老项目里的 error](example/3.compatibility-error.go)
1. 更多等待更新中...

输出普通信息和堆栈信息(string or array)
```go
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
    if e.Assert(err) { // 需要判断是否是自定义error, 否则无法输出堆栈信息.
        log.Println(e.Convert(err).ToStr())
    }
    // out:
    //2021/01/15 20:23:21 file:1.how.go, line:10, func:foo
    //file:1.how.go, line:15, func:main

    // (3)输出堆栈信息 by array
    if e.Assert(err) { // 需要判断是否是自定义error, 否则无法输出堆栈信息.
        log.Println(e.Convert(err).ToArr())
    }
    // out
    //2021/01/15 20:23:21 [file:1.how.go, line:10, func:foo file:1.how.go, line:15, func:main]
}
```

带自定义`code`的错误信息

```go
package main

import (
	"fmt"

	"github.com/yezihack/e"
)

// 如果使用带 code 的 error
func fooCode() error {
	return e.NewCode(400, "eoo error")
}

func main() {
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
```
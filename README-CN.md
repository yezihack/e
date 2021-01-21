![workflows](https://github.com/yezihack/e/workflows/Go/badge.svg)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/yezihack/e/Go)
[![](https://img.shields.io/badge/GoDoc-reference-green)](https://pkg.go.dev/github.com/yezihack/e)
[![GitHub license](https://img.shields.io/github/license/yezihack/e)](https://github.com/yezihack/e/blob/main/LICENSE)

[English](README.md)
# golang 优雅追踪堆栈错误包
> Golang tracks stack error package. 优雅追踪堆栈错误包

## 介绍(Introduction)
`github.com/yezihack/e` 项目是一个优雅地追踪你的堆栈信息.方便存储日志里.
而且还扩展了`error`包,自定义 `code,msg` 信息.

## 特色(Features)
> 优雅地追踪堆栈错误信息
1. 基于`github.com/pkg/errors`包进行封装
2. 支持 `code`, `msg` 自定义错误码和错误信息
3. 方便存储日志`json`文件
4. 堆栈信息以人性化展示

## 文档(Documentation)
[https://godoc.org/github.com/yezihack/e](https://godoc.org/github.com/yezihack/e)

## 安装(Install)
`go get github.com/yezihack/e`

## 简单使用(Use)
```go
package main
import (
	"github.com/yezihack/e"
    "log"
)
func foo() error {
	return e.New("foo")
}
func main() {
    err := foo()
    if err != nil { // 需要判断是否是自定义error, 否则无法输出堆栈信息.
        if e.Assert(err)  {
            log.Println(e.Convert(err).ToStr()) // 输出字符串形式
            log.Println(e.Convert(err).ToArr()) // 输出数组形式
        } else {
            log.Println(err) // 系统的 error
        }
    }
}
```

## 测试(Testing)
`go get github.com/smartystreets/goconvey`

`goconvey -port 8080`

## 实例(Example)
1. [基本用法](example/1.how_test.go)
1. [Code用法](example/2.code_test.go)
1. [兼容老项目里的 error](example/3.compatibility-error_test.go)
1. [获取 extra 的扩展错误](example/4.extra_test.go)
1. [gin中使用](example/5.gin_test.go)
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
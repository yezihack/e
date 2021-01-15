# golang 追踪堆栈错误包
> Golang tracks stack error package. 追踪堆栈错误包

## 输出堆栈信息
基于`github.com/pkg/errors`包进行封装,方便开发项目中使用. 

这个包本身是可以输出堆栈信息的,但是输出是原始的堆栈信息

存储在日志里查看不友好,而`github.com/yezihack/e`对堆栈信息进行封装了.

而且重新实现了系统自带的`error`接口,提供 `code` 自定义功能.

## 安装(install)
`go get github.com/yezihack/e`


## Example
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

```
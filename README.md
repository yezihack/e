[![workflows](https://github.com/yezihack/e/workflows/Go/badge.svg)](https://github.com/yezihack/e/actions?query=workflow%3AGo)
[![codecov](https://codecov.io/gh/yezihack/e/branch/main/graph/badge.svg?token=QZSV2DZM60)](https://codecov.io/gh/yezihack/e)
[![](https://img.shields.io/badge/GoDoc-reference-green)](https://pkg.go.dev/github.com/yezihack/e)
[![Go Report Card](https://goreportcard.com/badge/github.com/yezihack/e)](https://goreportcard.com/report/github.com/yezihack/e)
[![Gitter](https://img.shields.io/gitter/room/yezihack/e)](https://gitter.im/yezihack-e/community)

* [Golang graceful trace stack error packet](#golang-graceful-trace-stack-error-packet)
  * [Install](#install)
  * [Introduction](#introduction)
  * [Features](#features)
  * [Documentation](#documentation)
  * [Simple use](#simple-use)
  * [Testing](#testing)
  * [Example](#example)

[中文README](README-CN.md)
# Golang graceful trace stack error package
>Golang tracks stack error package


## Install
```
go get github.com/yezihack/e
```

## Introduction
`github.com/yezihack/e ` Project is a graceful way to track your stack information. 
Easy to store in log, It also extends the 'error' package, custom 'code, msg' information.

## Features
>Trace stack error messages gracefully

1. Based on` github.com/pkg/errors `Package
2. Support 'code', 'msg' custom error code and error information
3. Convenient storage of log 'JSON' files
4. The stack information is displayed humanized

## Documentation
[https://godoc.org/github.com/yezihack/e](https://godoc.org/github.com/yezihack/e)

## Simple use
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
    if err != nil { // you need to determine whether it is a custom error, otherwise you cannot output stack information
        if e.Assert(err)  {
            log.Println(e.Convert(err).ToStr()) // output string form
            log.Println(e.Convert(err).ToArr()) // output array form
        } else {
            log.Println(err) // system error
        }
    }
}
```
## Testing
```
go get github.com/smartystreets/goconvey
goconvey -port 8080
```


## Example
1. [basic usage](example/1.how_test.go)
1. [code usage](example/2.code_test.go)
1. [compatible with error in old projects](example/3.compatibility-error_test.go)
1. [get extension error of extra](example/4.extra_test.go)
1. [used in gin](example/5.gin_test.go))
1. More waiting to be updated

Output general information and stack information (string or array)
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
	// (1) general use
    err := foo("stack error")
    if err != nil {
        log.Println(err)
    }
    // out:
    // 2021/01/15 20:23:21 foo,stack error

    // (2)output stack information by string
    if e.Assert(err) { // you need to determine whether it is a custom error, otherwise you cannot output stack information.
        log.Println(e.Convert(err).ToStr())
    }
    // out:
    //2021/01/15 20:23:21 file:1.how.go, line:10, func:foo
    //file:1.how.go, line:15, func:main

    // (3) output stack information by array
    if e.Assert(err) { // you need to determine whether it is a custom error, otherwise you cannot output stack information.
        log.Println(e.Convert(err).ToArr())
    }
    // out
    //2021/01/15 20:23:21 [file:1.how.go, line:10, func:foo file:1.how.go, line:15, func:main]
}
```

Error message with custom 'code'

```go
package main

import (
	"fmt"

	"github.com/yezihack/e"
)

// if you use error with code
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
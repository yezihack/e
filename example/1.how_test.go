package main

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/yezihack/e"
)

func foo() error {
	return e.New("foo")
}
func TestFoo(t *testing.T) {
	err := foo()
	Convey("foo", t, func() {
		Convey("(1)普通使用", func() {
			So(err.Error(), ShouldEqual, "foo")
			// output: foo
		})
		Convey("(2)输出堆栈信息 by string", func() {
			// (2)输出堆栈信息 by string
			if err != nil && e.Assert(err) { // 需要判断是否是自定义error, 否则无法输出堆栈信息.
				str := e.Convert(err).ToStr()
				So(len(str), ShouldBeGreaterThan, 0)
				log.Println(str)
				// output:
				//2021/01/15 20:23:21 file:1.how.go, line:10, func:foo
				//file:1.how.go, line:15, func:main
			}
		})
		Convey("(3)输出堆栈信息 by array", func() {
			if e.Assert(err) { // 需要判断是否是自定义error, 否则无法输出堆栈信息.
				arr := e.Convert(err).ToArr()
				So(len(arr), ShouldBeGreaterThan, 0)
				log.Println(arr)
			}
			// output: 2021/01/15 20:23:21 [file:1.how.go, line:10, func:foo file:1.how.go, line:15, func:main]
		})
	})
}

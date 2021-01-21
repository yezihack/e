package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yezihack/e"
)

// 如果使用带 code 的 error
func fooCode() error {
	return e.NewCode(400, "eoo error")
}
func TestFooCode(t *testing.T) {
	err := fooCode()
	Convey("fooCode", t, func() {
		Convey("(1)输出错误码和错误信息", func() {
			if e.Assert(err) {
				e1 := e.Convert(err)
				So(e1.Code(), ShouldEqual, 400)
				So(e1.Msg(), ShouldEqual, "eoo error")
				fmt.Printf("code:%d, err:%s\n", e1.Code(), e1.Msg())
				//output:code:400, err:eoo error
			}
		})
		Convey("(2)输出错误码和错误堆栈信息", func() {
			if e.Assert(err) {
				e1 := e.Convert(err)
				So(e1.Code(), ShouldEqual, 400)
				So(len(e1.ToStr()), ShouldBeGreaterThan, 0)
				fmt.Printf("code:%d, err:%s\n", e1.Code(), e1.ToStr())
			}
			// output:
			//code:400, err:file:2.code.go, line:11, func:fooCode
			//file:2.code.go, line:15, func:main
		})
	})

}

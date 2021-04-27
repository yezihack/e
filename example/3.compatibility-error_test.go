package main

import (
	"errors"
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yezihack/e"
)

// 3. 兼容老项目里的 error

func oldFoo(x, y int) (err error) {
	if x < 0 || y < 0 {
		err = errors.New("x, y not less zero") // 原始的错误日志

		return
	}
	return nil
}
func MyFoo() error {
	err := oldFoo(-1, 10)
	if err != nil {
		err = e.WithStack(err) // 加入`github.com/yezihack/e`堆栈错误
		return err
	}
	return nil
}
func TestMyFoo(t *testing.T) {
	err := MyFoo()
	Convey("兼容已有的error", t, func() {
		if err != nil && e.Assert(err) {
			a := e.Convert(err).ToArr()
			So(len(a), ShouldBeGreaterThan, 0)
			So(a[0], ShouldContainSubstring, "func:MyFoo")
			So(a[1], ShouldContainSubstring, "func:TestMyFoo")
			log.Println(e.Convert(err).ToStr())
			//output:
			//2021/01/15 20:57:01 [file:3.compatibility-error.go, line:15, func:oldFoo file:3.compatibility-error.go, line:21, func:MyFoo file:3.compatibility-error.go, line:29, func:main]
		}
	})

}

package main

import (
	"log"
	"testing"

	"fmt"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yezihack/e"
)

// 添加扩展错误信息
func div(x, y int) (int, error) {
	if y == 0 {
		e1 := e.ErrorF("x:%d, y:%d", x, y)
		err := e.New("除数不能等于0", e1)
		return 0, err
	}
	return x / y, nil
}
func TestDiv(t *testing.T) {
	_, err := div(1, 0)
	Convey("添加扩展错误信息", t, func() {
		if e.Assert(err) {
			extra := e.Convert(err).Errs()
			So(len(extra), ShouldBeGreaterThan, 0)
			// 输出 extra 里的错误
			fmt.Printf("extra-error:%v\n", extra)
			// output: extra-error:[x:1, y:0]
			// 输出堆栈信息
			log.Println(e.Convert(err).ToStr())
			// output:
			// file:4.extra_test.go, line:17, func:div
			//file:4.extra_test.go, line:23, func:TestDiv
		}
	})

}

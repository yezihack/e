package e

import (
	"fmt"
	"testing"

	syserr "errors"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func foo() error {
	return errors.New("foo")
}
func getFoo() error {
	return foo()
}
func sysFoo() error {
	return syserr.New("sys foo")
}

func TestMarshalStack(t *testing.T) {

	Convey("MarshalStack", t, func() {
		Convey("MarshalStack-1", func() {
			err := getFoo()
			lst := MarshalStack(err)
			for _, item := range lst {
				fmt.Printf("file:%s,line:%s,func:%s\n", item.File, item.LineCode, item.FuncName)
			}
			So(len(lst), ShouldBeGreaterThan, 0)
			fmt.Println("MarshalStack-1 end")
		})
		Convey("MarshalStack-2", func() {
			err := sysFoo()
			lst := MarshalStack(err)
			So(len(lst), ShouldEqual, 0)
		})
	})
}

func Test_state_Width(t *testing.T) {
	Convey("state_Width", t, func() {
		s := state{}
		wid, ok := s.Width()
		ShouldBeFalse(ok)
		ShouldBeZeroValue(wid)
	})
}

func Test_state_Precision(t *testing.T) {
	Convey("state_Precision", t, func() {
		s := state{}
		wid, ok := s.Precision()
		ShouldBeFalse(ok)
		ShouldBeZeroValue(wid)
	})
}

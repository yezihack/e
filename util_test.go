package e

import (
	syserr "errors"
	"testing"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func sysFoo1() error {
	return syserr.New("sys foo")
}
func foo1() error {
	return errors.New("foo")
}

func TestToStr(t *testing.T) {
	Convey("ToStr", t, func() {
		Convey("ToStr-1", func() {
			So(ToStr(foo1()), ShouldContainSubstring, "TestToStr")
		})
		Convey("ToStr-2", func() {
			So(ToStr(sysFoo1()), ShouldEqual, "sys foo")
		})

	})
}

func TestToArr(t *testing.T) {
	Convey("ToArr", t, func() {
		Convey("ToArr-1", func() {
			err := foo1()
			arr := ToArr(err)
			So(len(arr), ShouldBeGreaterThan, 0)
			So(arr[0], ShouldContainSubstring, "TestToArr")
		})

		Convey("ToArr-2", func() {
			err := sysFoo1()
			arr := ToArr(err)
			So(len(arr), ShouldBeGreaterThan, 0)
			So(arr[0], ShouldContainSubstring, "sys foo")
		})
	})
}

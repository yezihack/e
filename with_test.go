package e

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestErrorF(t *testing.T) {
	Convey("ErrorF", t, func() {
		err := func() error {
			return ErrorF("my error:%s", "hello")
		}
		So(err(), ShouldNotBeNil)
	})
}

func TestNew(t *testing.T) {
	Convey("New", t, func() {
		err := func() error {
			return New("hello e")
		}
		So(err(), ShouldNotBeNil)
	})
}

func TestWithMessage(t *testing.T) {
	Convey("WithMessage", t, func() {
		err := func() error {
			return WithMessage(errors.New("hello"), "msg")
		}
		So(err(), ShouldNotBeNil)
	})
}

func TestWithMessageF(t *testing.T) {
	Convey("WithMessageF", t, func() {
		err := func() error {
			return WithMessageF(errors.New("hello"), "(%s)", "msg")
		}
		So(err(), ShouldNotBeNil)
	})
}

func TestWithStack(t *testing.T) {
	Convey("WithMessageF", t, func() {
		err := func() error {
			return WithStack(errors.New("hello"))
		}
		So(err(), ShouldNotBeNil)
	})
}

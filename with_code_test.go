package e

import (
	"testing"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestErrorCodeF(t *testing.T) {
	Convey("ErrorCodeF", t, func() {
		err := func() error {
			return ErrorCodeF(100, "my error:%s", "hello")
		}
		ShouldNotBeNil(err())
	})
}

func TestNewCode(t *testing.T) {
	Convey("NewCode", t, func() {
		err := func() error {
			return NewCode(101, "hello e")
		}
		ShouldNotBeNil(err())
	})
}

func TestWithCodeMessage(t *testing.T) {
	Convey("WithCodeMessage", t, func() {
		err := func() error {
			return WithCodeMessage(102, errors.New("hello"), "msg")
		}
		So(err(), ShouldNotBeNil)
	})
}

func TestWithCodeMessageF(t *testing.T) {
	Convey("WithCodeMessageF", t, func() {
		err := func() error {
			return WithCodeMessageF(103, errors.New("hello"), "(%s)", "msg")
		}
		So(err(), ShouldNotBeNil)
	})
}

func TestWithCodeStack(t *testing.T) {
	Convey("WithCodeStack", t, func() {
		err := func() error {
			return WithCodeStack(104, errors.New("hello"))
		}
		So(err(), ShouldNotBeNil)
	})
}

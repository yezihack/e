package e

import (
	"testing"

	"github.com/pkg/errors"

	. "github.com/smartystreets/goconvey/convey"
)

func MyErr(code int, err error, extra ...error) error {
	return &StackError{
		code:      code,
		err:       err,
		extraErrs: extra,
	}
}

func TestStackError_Code(t *testing.T) {
	Convey("StackError_Code", t, func() {
		err := MyErr(200, nil)
		if e := err.(*StackError); e != nil {
			So(e.Code(), ShouldEqual, 200)
		}
	})
}

func TestStackError_Err(t *testing.T) {
	Convey("StackError_Err", t, func() {
		err := MyErr(200, errors.New("foo"))
		if e := err.(*StackError); e != nil {
			So(e.Err(), ShouldBeError, errors.New("foo"))
		}
	})
}

func TestStackError_Error(t *testing.T) {
	Convey("StackError_Error", t, func() {
		err := MyErr(200, errors.New("foo"))
		if e := err.(*StackError); e != nil {
			So(e.Error(), ShouldEqual, "foo")
		}
	})
}

func TestStackError_Errs(t *testing.T) {
	Convey("StackError_Errs", t, func() {
		err := MyErr(200, errors.New("foo"), errors.New("f1"), errors.New("f2"))
		if e := err.(*StackError); e != nil {
			So(e.Errs()[0], ShouldBeError, errors.New("f1"))
			So(e.Errs()[1], ShouldBeError, errors.New("f2"))
		}
	})
}

func TestStackError_ExistExtra(t *testing.T) {
	Convey("StackError_ExistExtra", t, func() {
		err := MyErr(200, errors.New("foo"), errors.New("f1"), errors.New("f2"))
		if e := err.(*StackError); e != nil {
			So(e.ExistExtra(), ShouldBeTrue)
		}
		err2 := MyErr(200, errors.New("foo"))
		if e := err2.(*StackError); e != nil {
			So(e.ExistExtra(), ShouldBeFalse)
		}
	})
}

func TestStackError_Msg(t *testing.T) {
	Convey("StackError_Msg", t, func() {
		err := MyErr(200, errors.New("foo"), errors.New("f1"), errors.New("f2"))
		if e := err.(*StackError); e != nil {
			So(e.Msg(), ShouldEqual, "foo")
		}
	})
}

func TestStackError_ToArr(t *testing.T) {
	Convey("StackError_ToArr", t, func() {
		err := MyErr(200, errors.New("foo"))
		if e := err.(*StackError); e != nil {
			So(len(e.ToArr()), ShouldBeGreaterThan, 0)
		}
	})
}

func TestStackError_ToStr(t *testing.T) {
	Convey("StackError_ToStr", t, func() {
		err := MyErr(200, errors.New("foo"))
		if e := err.(*StackError); e != nil {
			So(e.ToStr(), ShouldContainSubstring, "TestStackError_ToStr")
		}
	})
}

func TestStackError_ToStrByExtra(t *testing.T) {
	Convey("StackError_ToStrByExtra", t, func() {
		err := MyErr(200, errors.New("foo"), errors.New("foo1"), errors.New("foo2"))
		if e := err.(*StackError); e != nil {
			So(e.ToStrByExtra(), ShouldContainSubstring, "foo1")
			So(e.ToStrByExtra(), ShouldContainSubstring, "foo2")
		}
	})
}

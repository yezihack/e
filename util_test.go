package e

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func foo1() error {
	return errors.New("foo")
}
func TestToStr(t *testing.T) {
	Convey("ToStr", t, func() {
		err := foo1()
		e1 := ToStr(err)
		fmt.Println(e1)
	})
}

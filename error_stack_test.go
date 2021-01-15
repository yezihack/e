package e

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func foo() error {
	return errors.New("foo")
}

func TestMarshalStack(t *testing.T) {
	Convey("MarshalStack", t, func() {
		err := foo()
		lst := MarshalStack(err)
		for _, item := range lst {
			fmt.Printf("file:%s,line:%s,func:%s\n", item.File, item.LineCode, item.FuncName)
		}
	})
}

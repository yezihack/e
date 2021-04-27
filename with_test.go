package e

import (
	"fmt"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)
func withErr()error {
	return With().Msg("not found id")
}

func TestWith(t *testing.T) {
	Convey("with", t, func() {
		err := withErr()
		So(Convert(err).Msg(), ShouldEqual, "not found id")
		So(Convert(err).Code(), ShouldEqual, 0)
		So(Convert(err).Error(),ShouldEqual, "")
		So(Convert(err).ToStr(), ShouldNotBeEmpty)
	})
}

func Test_with_Err(t *testing.T) {
	Convey("with", t, func() {
		err := With().Err(errors.New("uid is not exist"))
		So(Convert(err).Msg(), ShouldEqual, "")
		So(Convert(err).Code(), ShouldEqual, 0)
		So(Convert(err).Error(),ShouldEqual, "uid is not exist")
		So(Convert(err).ToStr(), ShouldNotBeEmpty)
	})
}

func Test_with_ErrMsg(t *testing.T) {
	Convey("with", t, func() {
		err := With().ErrMsg(errors.New("uid is not exist"), "github")
		So(Convert(err).Msg(), ShouldEqual, "github")
		So(Convert(err).Code(), ShouldEqual, 0)
		So(Convert(err).Error(),ShouldEqual, "uid is not exist")
		fmt.Printf("%+v", Convert(err).Err())
		So(Convert(err).ToStr(), ShouldNotBeEmpty)
	})
}

func Test_with_Msg(t *testing.T) {
	type args struct {
		s      string
		extras []error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &with{}
			if err := w.Msg(tt.args.s, tt.args.extras...); (err != nil) != tt.wantErr {
				t.Errorf("Msg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
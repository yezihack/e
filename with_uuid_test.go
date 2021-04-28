package e

import (
	"reflect"
	"testing"
)

func TestWithUUID(t *testing.T) {
	tests := []struct {
		name string
		want WithUUIDer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithUUID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_withUUID_Code(t *testing.T) {
	type args struct {
		code   int
		uuid   string
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
			w := &withUUID{}
			if err := w.Code(tt.args.code, tt.args.uuid, tt.args.extras...); (err != nil) != tt.wantErr {
				t.Errorf("Code() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_withUUID_Err(t *testing.T) {
	type args struct {
		code   int
		uuid   string
		err    error
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
			w := &withUUID{}
			if err := w.Err(tt.args.code, tt.args.uuid, tt.args.err, tt.args.extras...); (err != nil) != tt.wantErr {
				t.Errorf("Err() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_withUUID_ErrMsg(t *testing.T) {
	type args struct {
		code   int
		uuid   string
		err    error
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
			w := &withUUID{}
			if err := w.ErrMsg(tt.args.code, tt.args.uuid, tt.args.err, tt.args.s, tt.args.extras...); (err != nil) != tt.wantErr {
				t.Errorf("ErrMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_withUUID_Msg(t *testing.T) {
	type args struct {
		code   int
		uuid   string
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
			w := &withUUID{}
			if err := w.Msg(tt.args.code, tt.args.uuid, tt.args.s, tt.args.extras...); (err != nil) != tt.wantErr {
				t.Errorf("Msg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_withUUID_UUID(t *testing.T) {
	type args struct {
		uuid   string
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
			w := &withUUID{}
			if err := w.UUID(tt.args.uuid, tt.args.extras...); (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

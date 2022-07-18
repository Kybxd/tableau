package log

import (
	"testing"
)

func TestDebug(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				args: []interface{}{"xxx", 1, true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.args...)
		})
	}
}

func TestInfow(t *testing.T) {
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				msg:           "infow test",
				keysAndValues: []interface{}{"xxx", 1, "key2", true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infow(tt.args.msg, tt.args.keysAndValues...)
			With("key1", 123).Debug(tt.args.keysAndValues...)
		})
	}
}

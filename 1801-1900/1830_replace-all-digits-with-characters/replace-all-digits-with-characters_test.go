package main

import "testing"

func Test_replaceDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"a1c1e1"}, "abcdef"},
		{"2", args{"a1b2c3d4e"}, "abbdcfdhe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceDigits(tt.args.s); got != tt.want {
				t.Errorf("replaceDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

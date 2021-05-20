package main

import "testing"

func Test_canTransform(t *testing.T) {
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"RXXLRXRXL", "XRLXXRRLX"}, true},
		{"2", args{"x", "L"}, false},
		{"3", args{"LLR", "RRL"}, false},
		{"4", args{"XL", "LX"}, true},
		{"5", args{"XLLR", "LXLX"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canTransform(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("canTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

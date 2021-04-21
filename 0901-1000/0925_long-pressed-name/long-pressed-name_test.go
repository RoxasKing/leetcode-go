package main

import "testing"

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"alex", "aaleex"}, true},
		{"2", args{"saeed", "ssaaedd"}, false},
		{"3", args{"leelee", "lleeelee"}, true},
		{"4", args{"laiden", "laiden"}, true},
		{"5", args{"pyplrz", "ppyypllr"}, false},
		{"6", args{"vtkgn", "vttkgnn"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}

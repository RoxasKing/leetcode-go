package main

import (
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"ab#c", "ad#c"}, true},
		{"2", args{"ab##", "c#d#"}, true},
		{"3", args{"a##c", "#a#c"}, true},
		{"4", args{"a#c", "b"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backspaceCompare2(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"ab#c", "ad#c"}, true},
		{"2", args{"ab##", "c#d#"}, true},
		{"3", args{"a##c", "#a#c"}, true},
		{"4", args{"a#c", "b"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare2(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_compressString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"aabcccccaaa"}, "a2b1c5a3"},
		{"2", args{"abbccd"}, "abbccd"},
		{"3", args{"a"}, "a"},
		{"4", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.S); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}

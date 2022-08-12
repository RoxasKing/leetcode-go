package main

import "testing"

func Test_reformat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"a0b1c2"}, "0a1b2c"},
		{"2", args{"leetcode"}, ""},
		{"3", args{"1229857369"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformat(tt.args.s); got != tt.want {
				t.Errorf("reformat() = %v, want %v", got, tt.want)
			}
		})
	}
}

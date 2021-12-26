package main

import "testing"

func Test_repeatedStringMatch(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abcd", "cdabcdab"}, 3},
		{"2", args{"a", "aa"}, 2},
		{"3", args{"a", "a"}, 1},
		{"4", args{"abc", "wxyz"}, -1},
		{"5", args{"aaaaaaaaaaaaaaaaaaaaaab", "ba"}, 2},
		{"6", args{"abcbc", "cabcbca"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("repeatedStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

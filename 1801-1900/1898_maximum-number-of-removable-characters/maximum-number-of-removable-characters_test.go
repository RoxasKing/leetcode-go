package main

import "testing"

func Test_maximumRemovals(t *testing.T) {
	type args struct {
		s         string
		p         string
		removable []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abcacb", "ab", []int{3, 1, 0}}, 2},
		{"2", args{"abcbddddd", "abcd", []int{3, 2, 1, 4, 5, 6}}, 1},
		{"3", args{"abcab", "abc", []int{0, 1, 2, 3, 4}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRemovals(tt.args.s, tt.args.p, tt.args.removable); got != tt.want {
				t.Errorf("maximumRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}

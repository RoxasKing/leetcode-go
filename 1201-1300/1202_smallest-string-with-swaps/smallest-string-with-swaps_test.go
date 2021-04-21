package main

import "testing"

func Test_smallestStringWithSwaps(t *testing.T) {
	type args struct {
		s     string
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"dcab", [][]int{{0, 3}, {1, 2}}}, "bacd"},
		{"2", args{"dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}}, "abcd"},
		{"3", args{"cba", [][]int{{0, 1}, {1, 2}}}, "abc"},
		{"4", args{"c", [][]int{}}, "c"},
		{"5", args{"cb", [][]int{{0, 1}}}, "bc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestStringWithSwaps(tt.args.s, tt.args.pairs); got != tt.want {
				t.Errorf("smallestStringWithSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

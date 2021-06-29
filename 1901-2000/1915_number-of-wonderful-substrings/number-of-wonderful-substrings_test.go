package main

import "testing"

func Test_wonderfulSubstrings(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{"aba"}, 4},
		{"2", args{"aabb"}, 9},
		{"3", args{"he"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wonderfulSubstrings(tt.args.word); got != tt.want {
				t.Errorf("wonderfulSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

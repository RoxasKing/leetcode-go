package main

import "testing"

func Test_countGoodSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"xyzzaz"}, 1},
		{"2", args{"aababcabc"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countGoodSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

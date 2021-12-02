package main

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"leetcode"}, 2},
		{"2", args{"abbcccddddeeeeedcba"}, 5},
		{"3", args{"triplepillooooow"}, 5},
		{"4", args{"hooraaaaaaaaaaay"}, 11},
		{"5", args{"tourist"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.s); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

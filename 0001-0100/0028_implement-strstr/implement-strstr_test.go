package main

import (
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"hello", "ll"}, 2},
		{"2", args{"aaaaa", "bba"}, -1},
		{"3", args{"", ""}, 0},
		{"4", args{"hello", "o"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStr2(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"hello", "ll"}, 2},
		{"2", args{"aaaaa", "bba"}, -1},
		{"3", args{"", ""}, 0},
		{"4", args{"aabaacaabaab", "aabaab"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr2(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr2() = %v, want %v", got, tt.want)
			}
		})
	}
}

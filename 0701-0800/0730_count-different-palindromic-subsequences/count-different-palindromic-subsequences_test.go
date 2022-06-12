package main

import "testing"

func Test_countPalindromicSubsequences(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"bccb"}, 6},
		{"2", args{"abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba"}, 104860361},
		{"3", args{"ab"}, 2},
		{"4", args{"aaa"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPalindromicSubsequences(tt.args.S); got != tt.want {
				t.Errorf("countPalindromicSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

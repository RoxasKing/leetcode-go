package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"cacb", "cbba"}, 5},
		{"2", args{"ab", "ab"}, 3},
		{"3", args{"aa", "bb"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

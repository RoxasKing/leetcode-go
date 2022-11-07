package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"lc", "cl", "gg"}}, 6},
		{"2", args{[]string{"ab", "ty", "yt", "lc", "cl", "ab"}}, 8},
		{"3", args{[]string{"cc", "ll", "xx"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.words); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

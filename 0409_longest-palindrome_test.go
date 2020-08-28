package main

import "testing"

func Test_longestPalindrome0409(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"abccccdd"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome0409(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome0409() = %v, want %v", got, tt.want)
			}
		})
	}
}

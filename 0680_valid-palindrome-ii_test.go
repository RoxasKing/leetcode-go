package main

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"aba"}, true},
		{"2", args{"abca"}, true},
		{"3", args{"abcda"}, false},
		{"4", args{"a"}, true},
		{"5", args{"aa"}, true},
		{"6", args{"eeccccbebaeeabebccceea"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_shortestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"aacecaaa"}, "aaacecaaa"},
		{"2", args{"abcd"}, "dcbabcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("shortestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortestPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"aacecaaa"}, "aaacecaaa"},
		{"2", args{"abcd"}, "dcbabcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("shortestPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

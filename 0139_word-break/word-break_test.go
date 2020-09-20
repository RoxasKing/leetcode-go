package main

import (
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"a", []string{"a"}}, true},
		{"2", args{"leetcode", []string{"leet", "code"}}, true},
		{"3", args{"applepenapple", []string{"apple", "pen"}}, true},
		{"4", args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

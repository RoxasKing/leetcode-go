package main

import "testing"

func Test_countConsistentStrings(t *testing.T) {
	type args struct {
		allowed string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"ab", []string{"ad", "bd", "aaab", "baa", "badab"}}, 2},
		{"2", args{"abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}}, 7},
		{"3", args{"cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countConsistentStrings(tt.args.allowed, tt.args.words); got != tt.want {
				t.Errorf("countConsistentStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

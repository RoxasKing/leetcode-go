package main

import (
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"}, []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"2", args{"AAAAAAAAAAAA"}, []string{"AAAAAAAAAA"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedDnaSequences(tt.args.s); !equal(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(s1, s2 []string) bool {
	src := make(map[string]bool)
	dst := make(map[string]bool)
	for _, s := range s1 {
		src[s] = true
	}
	for _, s := range s2 {
		dst[s] = true
	}
	for _, s := range s1 {
		if !dst[s] {
			return false
		}
	}
	for _, s := range s2 {
		if !src[s] {
			return false
		}
	}
	return true
}

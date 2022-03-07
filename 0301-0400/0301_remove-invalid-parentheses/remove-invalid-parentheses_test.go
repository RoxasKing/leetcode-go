package main

import "testing"

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"()())()"}, []string{"()()()", "(())()"}},
		{"2", args{"(a)())()"}, []string{"(a)()()", "(a())()"}},
		{"3", args{")(f"}, []string{"f"}},
		{"4", args{"))((f"}, []string{"f"}},
		{"5", args{"((()())()"}, []string{"(()())()", "((()()))", "((()))()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeInvalidParentheses(tt.args.s); !equal(got, tt.want) {
				t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
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

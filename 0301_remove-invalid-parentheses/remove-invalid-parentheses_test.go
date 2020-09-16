package main

import (
	"testing"
)

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeInvalidParentheses(tt.args.s)
			src := make(map[string]bool)
			dst := make(map[string]bool)
			for _, s := range got {
				src[s] = true
			}
			for _, s := range tt.want {
				dst[s] = true
			}
			for _, s := range got {
				if !dst[s] {
					t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
					return
				}
			}
			for _, s := range tt.want {
				if !src[s] {
					t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func Test_removeInvalidParentheses2(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeInvalidParentheses2(tt.args.s)
			src := make(map[string]bool)
			dst := make(map[string]bool)
			for _, s := range got {
				src[s] = true
			}
			for _, s := range tt.want {
				dst[s] = true
			}
			for _, s := range got {
				if !dst[s] {
					t.Errorf("removeInvalidParentheses2() = %v, want %v", got, tt.want)
					return
				}
			}
			for _, s := range tt.want {
				if !src[s] {
					t.Errorf("removeInvalidParentheses2() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

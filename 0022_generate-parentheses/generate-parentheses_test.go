package main

import (
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{3},
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.args.n)
			memo1, memo2 := make(map[string]bool), make(map[string]bool)
			for _, str := range got {
				memo1[str] = true
			}
			for _, str := range tt.want {
				memo2[str] = true
			}
			for _, str := range got {
				if !memo2[str] {
					t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
				}
			}
			for _, str := range tt.want {
				if !memo1[str] {
					t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_generateParenthesis2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{3},
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.args.n)
			memo1, memo2 := make(map[string]bool), make(map[string]bool)
			for _, str := range got {
				memo1[str] = true
			}
			for _, str := range tt.want {
				memo2[str] = true
			}
			for _, str := range got {
				if !memo2[str] {
					t.Errorf("generateParenthesis2() = %v, want %v", got, tt.want)
				}
			}
			for _, str := range tt.want {
				if !memo1[str] {
					t.Errorf("generateParenthesis2() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

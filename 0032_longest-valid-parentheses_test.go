package leetcode

import (
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{")(()())"}, 6},
		{"", args{")()())"}, 4},
		{"", args{"(()"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidParentheses2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"(()"}, 2},
		{"", args{")()())"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses2(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidParentheses3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"(()"}, 2},
		{"", args{")()())"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses3(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses3() = %v, want %v", got, tt.want)
			}
		})
	}
}

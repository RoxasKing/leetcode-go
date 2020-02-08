package leetcode

import (
	"testing"
)

func Test_isMatch_0044(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"acdcb", "a*c?b"}, false},
		{"", args{"adceb", "*a*b"}, true},
		{"", args{"cb", "?a"}, false},
		{"", args{"aa", "a"}, false},
		{"", args{"aa", "*"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch0044(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch_0044() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatch_0044_2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"acdcb", "a*c?b"}, false},
		{"", args{"adceb", "*a*b"}, true},
		{"", args{"cb", "?a"}, false},
		{"", args{"aa", "a"}, false},
		{"", args{"aa", "*"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch00442(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch_0044_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_isMatchII(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"acdcb", "a*c?b"}, false},
		{"2", args{"adceb", "*a*b"}, true},
		{"3", args{"cb", "?a"}, false},
		{"4", args{"aa", "a"}, false},
		{"5", args{"aa", "*"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatchII(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch_0044() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatchII2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"acdcb", "a*c?b"}, false},
		{"2", args{"adceb", "*a*b"}, true},
		{"3", args{"cb", "?a"}, false},
		{"4", args{"aa", "a"}, false},
		{"5", args{"aa", "*"}, true},
		{"6", args{"ab", "*a"}, false},
		{"7", args{"a", "*a"}, true},
		{"8", args{"aaa", "**a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatchII2(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch_0044_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package codinginterviews

import (
	"testing"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"aa", "a"}, false},
		{"2", args{"aa", "a*"}, true},
		{"3", args{"ab", ".*"}, true},
		{"4", args{"aab", "c*a*b"}, true},
		{"5", args{"mississippi", "mis*is*p*."}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatch2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"aa", "a"}, false},
		{"2", args{"aa", "a*"}, true},
		{"3", args{"ab", ".*"}, true},
		{"4", args{"aab", "c*a*b"}, true},
		{"5", args{"mississippi", "mis*is*p*."}, false},
		{"6", args{"ab", ".*c"}, false},
		{"7", args{"a", ".*..a*"}, false},
		{"8", args{"", ".*"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch2(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch2() = %v, want %v", got, tt.want)
			}
		})
	}
}

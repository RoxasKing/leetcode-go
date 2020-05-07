package leetcode

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"abcabcbb"}, 3},
		{"", args{"bbbbbbbb"}, 1},
		{"", args{"pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstring2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"abcabcbb"}, 3},
		{"", args{"bbbbbbbb"}, 1},
		{"", args{"pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{"abcabcbb"}, "abc"},
		{"test 2", args{"bbbbbbbb"}, "b"},
		{"test 3", args{"pwwkew"}, "wke"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

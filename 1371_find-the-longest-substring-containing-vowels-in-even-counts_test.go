package leetcode

import "testing"

func Test_findTheLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"eleetminicoworoep"}, 13},
		{"2", args{"leetcodeisgreat"}, 5},
		{"3", args{"bcbcbc"}, 6},
		{"4", args{"id"}, 1},
		{"5", args{"ixzhsdka"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("findTheLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

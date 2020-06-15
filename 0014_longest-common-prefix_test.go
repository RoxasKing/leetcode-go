package leetcode

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]string{"aa", "a"}}, "a"},
		{"2", args{[]string{"", ""}}, ""},
		{"3", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"4", args{[]string{"dog", "racecar", "car"}}, ""},
		{"5", args{[]string{"c", "c"}}, "c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

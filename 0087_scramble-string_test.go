package leetcode

import "testing"

func Test_isScramble(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"rgeat", "great"}, true},
		{"", args{"abcde", "caebd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}

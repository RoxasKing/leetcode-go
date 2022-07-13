package main

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"aabcc", "dbbca", "aadbbcbcac"}, true},
		{"2", args{"aabcc", "dbbca", "aadbbbaccc"}, false},
		{"3", args{"", "", ""}, true},
		{"4", args{"a", "", "c"}, false},
		{"5", args{"a", "b", "ac"}, false},
		{"6", args{"a", "b", "ab"}, true},
		{"7", args{"db", "b", "cbb"}, false},
		{"8", args{"", "ab", "ab"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_makeFancyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"leeetcode"}, "leetcode"},
		{"2", args{"aaabaaaa"}, "aabaa"},
		{"3", args{"aab"}, "aab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeFancyString(tt.args.s); got != tt.want {
				t.Errorf("makeFancyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

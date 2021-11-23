package main

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"ab", "ba"}, true},
		{"2", args{"ab", "ab"}, false},
		{"3", args{"aa", "aa"}, true},
		{"4", args{"aaaaaaabc", "aaaaaaacb"}, true},
		{"5", args{"abcaa", "abcbb"}, false},
		{"6", args{"ab", "babbb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

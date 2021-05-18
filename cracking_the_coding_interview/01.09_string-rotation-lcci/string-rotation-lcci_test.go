package main

import "testing"

func Test_isFlipedString(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"waterbottle", "erbottlewat"}, true},
		{"2", args{"aa", "aba"}, false},
		{"3", args{"", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFlipedString(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isFlipedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

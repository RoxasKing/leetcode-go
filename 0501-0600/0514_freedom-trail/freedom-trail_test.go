package main

import (
	"testing"
)

func Test_findRotateSteps(t *testing.T) {
	type args struct {
		ring string
		key  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"godding", "gd"}, 4},
		{"2", args{"goddings", "gds"}, 8},
		{"3", args{"ababcab", "acbaacba"}, 17},
		{"4", args{"caotmcaataijjxi", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx"}, 137},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotateSteps(tt.args.ring, tt.args.key); got != tt.want {
				t.Errorf("findRotateSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

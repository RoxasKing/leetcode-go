package main

import (
	"testing"
)

func Test_maxBuilding(t *testing.T) {
	type args struct {
		n            int
		restrictions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, [][]int{{2, 1}, {4, 1}}}, 2},
		{"2", args{6, [][]int{}}, 5},
		{"3", args{10, [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxBuilding(tt.args.n, tt.args.restrictions); got != tt.want {
				t.Errorf("maxBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}

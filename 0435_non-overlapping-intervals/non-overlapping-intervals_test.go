package main

import (
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}}, 1},
		{"2", args{[][]int{{1, 2}, {1, 2}, {1, 2}}}, 2},
		{"3", args{[][]int{{1, 2}, {2, 3}}}, 0},
		{"4", args{[][]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

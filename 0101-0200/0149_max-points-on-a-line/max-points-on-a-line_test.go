package main

import (
	"testing"
)

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 1}, {2, 2}, {3, 3}}}, 3},
		{"2", args{[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}}, 4},
		{"3", args{[][]int{{0, 0}, {94911150, 94911151}, {94911151, 94911152}}}, 2},
		{"4", args{[][]int{{4, 0}, {4, -1}, {4, 5}}}, 3},
		{"5", args{[][]int{{1, 1}, {1, 1}, {1, 1}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

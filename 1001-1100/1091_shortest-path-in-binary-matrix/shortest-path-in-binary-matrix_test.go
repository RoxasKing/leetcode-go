package main

import "testing"

func Test_shortestPathBinaryMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{0, 1}, {1, 0}}}, 2},
		{"2", args{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}}, 4},
		{"3", args{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathBinaryMatrix(tt.args.grid); got != tt.want {
				t.Errorf("shortestPathBinaryMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

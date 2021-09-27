package main

import "testing"

func Test_shortestPath(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}}, 1}, 6},
		{"2", args{[][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}, 1}, -1},
		{"3", args{[][]int{{0}}, 1}, 0},
		{"4", args{[][]int{{0, 0}}, 1}, 1},
		{"5", args{[][]int{{0}, {0}}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPath(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("shortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

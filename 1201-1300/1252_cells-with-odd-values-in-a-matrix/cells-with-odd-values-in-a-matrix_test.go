package main

import "testing"

func Test_oddCells(t *testing.T) {
	type args struct {
		m       int
		n       int
		indices [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 3, [][]int{{0, 1}, {1, 1}}}, 6},
		{"2", args{2, 2, [][]int{{1, 1}, {0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddCells(tt.args.m, tt.args.n, tt.args.indices); got != tt.want {
				t.Errorf("oddCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

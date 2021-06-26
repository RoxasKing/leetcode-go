package main

import "testing"

func Test_slidingPuzzle(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 2, 3}, {4, 0, 5}}}, 1},
		{"2", args{[][]int{{1, 2, 3}, {5, 4, 0}}}, -1},
		{"3", args{[][]int{{4, 1, 2}, {5, 0, 3}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingPuzzle(tt.args.board); got != tt.want {
				t.Errorf("slidingPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}

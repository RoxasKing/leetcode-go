package main

import "testing"

func Test_movesToChessboard(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}}}, 2},
		{"2", args{[][]int{{0, 1}, {1, 0}}}, 0},
		{"3", args{[][]int{{1, 0}, {1, 0}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movesToChessboard(tt.args.board); got != tt.want {
				t.Errorf("movesToChessboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

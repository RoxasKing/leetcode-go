package main

import "testing"

func Test_largestMagicSquare(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}}}, 3},
		{"2", args{[][]int{{5, 1, 3, 1}, {9, 3, 3, 1}, {1, 3, 3, 8}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestMagicSquare(tt.args.grid); got != tt.want {
				t.Errorf("largestMagicSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

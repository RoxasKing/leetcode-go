package main

import "testing"

func Test_uniquePathsIII(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}}, 2},
		{"2", args{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}}, 4},
		{"3", args{[][]int{{0, 1}, {2, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsIII(tt.args.grid); got != tt.want {
				t.Errorf("uniquePathsIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

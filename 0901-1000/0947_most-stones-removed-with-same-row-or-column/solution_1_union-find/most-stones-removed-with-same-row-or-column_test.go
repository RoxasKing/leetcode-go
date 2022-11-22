package main

import "testing"

func Test_removeStones(t *testing.T) {
	type args struct {
		stones [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}}, 5},
		{"2", args{[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}}, 3},
		{"3", args{[][]int{{0, 0}}}, 0},
		{"4", args{[][]int{{0, 1}, {1, 0}, {1, 1}}}, 2},
		{"5", args{[][]int{{0, 1}, {1, 2}, {1, 3}, {3, 3}, {2, 3}, {0, 2}}}, 5},
		{"6", args{[][]int{{0, 1}, {0, 2}, {4, 3}, {2, 4}, {0, 3}, {1, 1}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeStones(tt.args.stones); got != tt.want {
				t.Errorf("removeStones() = %v, want %v", got, tt.want)
			}
		})
	}
}

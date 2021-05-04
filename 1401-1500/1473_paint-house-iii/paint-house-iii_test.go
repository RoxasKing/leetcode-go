package main

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		houses []int
		cost   [][]int
		m      int
		n      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{0, 0, 0, 0, 0},
				[][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}},
				5, 2, 3,
			},
			9,
		},
		{
			"2",
			args{
				[]int{0, 2, 1, 2, 0},
				[][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}},
				5, 2, 3,
			},
			11,
		},
		{
			"3",
			args{
				[]int{0, 0, 0, 0, 0},
				[][]int{{1, 10}, {10, 1}, {1, 10}, {10, 1}, {1, 10}},
				5, 2, 5,
			},
			5,
		},
		{
			"4",
			args{
				[]int{3, 1, 2, 3},
				[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
				4, 3, 3,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.houses, tt.args.cost, tt.args.m, tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

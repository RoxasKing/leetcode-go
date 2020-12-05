package main

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}}, 2},
		{"2", args{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}}, 4},
		{"3", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}}, 2},
		{"4", args{[][]int{{1, 2}}}, 1},
		{"5", args{[][]int{{2, 3}, {2, 3}}}, 1},
		{"6", args{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}, {12, 18}}}, 2},
		{"7", args{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}, {12, 18}, {1, 15}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}

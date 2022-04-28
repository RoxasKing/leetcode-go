package main

import "testing"

func Test_minCostConnectPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}}, 20},
		{"2", args{[][]int{{3, 12}, {-2, 5}, {-4, 1}}}, 18},
		{"3", args{[][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}}, 4},
		{"4", args{[][]int{{-1000000, -1000000}, {1000000, 1000000}}}, 4000000},
		{"5", args{[][]int{{0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPoints(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_electricCarPlan(t *testing.T) {
	type args struct {
		paths  [][]int
		cnt    int
		start  int
		end    int
		charge []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[][]int{{1, 3, 3}, {3, 2, 1}, {2, 1, 3}, {0, 1, 4}, {3, 0, 5}},
				6,
				1,
				0,
				[]int{2, 10, 4, 1},
			},
			43,
		},
		{
			"2",
			args{
				[][]int{{0, 4, 2}, {4, 3, 5}, {3, 0, 5}, {0, 1, 5}, {3, 2, 4}, {1, 2, 8}},
				8,
				0,
				2,
				[]int{4, 1, 1, 3, 2},
			},
			38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := electricCarPlan(tt.args.paths, tt.args.cnt, tt.args.start, tt.args.end, tt.args.charge); got != tt.want {
				t.Errorf("electricCarPlan() = %v, want %v", got, tt.want)
			}
		})
	}
}

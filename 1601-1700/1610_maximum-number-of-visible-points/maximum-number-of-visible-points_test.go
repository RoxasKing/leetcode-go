package main

import "testing"

func Test_visiblePoints(t *testing.T) {
	type args struct {
		points   [][]int
		angle    int
		location []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[][]int{{2, 1}, {2, 2}, {3, 3}},
				90,
				[]int{1, 1},
			},
			3,
		},
		{
			"2",
			args{
				[][]int{{2, 1}, {2, 2}, {3, 4}, {1, 1}},
				90,
				[]int{1, 1},
			},
			4,
		},
		{
			"3",
			args{
				[][]int{{1, 0}, {2, 1}},
				13,
				[]int{1, 1},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := visiblePoints(tt.args.points, tt.args.angle, tt.args.location); got != tt.want {
				t.Errorf("visiblePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

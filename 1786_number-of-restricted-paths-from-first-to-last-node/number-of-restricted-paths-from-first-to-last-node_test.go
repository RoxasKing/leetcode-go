package main

import "testing"

func Test_countRestrictedPaths(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				5,
				[][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}},
			},
			3,
		},
		{
			"2",
			args{
				7,
				[][]int{{1, 3, 1}, {4, 1, 2}, {7, 3, 4}, {2, 5, 3}, {5, 6, 1}, {6, 7, 2}, {7, 5, 3}, {2, 6, 4}},
			},
			1,
		},
		{
			"3",
			args{
				6,
				[][]int{{6, 2, 9}, {2, 1, 7}, {6, 5, 10}, {4, 3, 1}, {3, 1, 8}, {4, 6, 4}, {5, 1, 7}, {1, 4, 7}, {4, 5, 3}, {3, 6, 6}, {5, 3, 9}, {1, 6, 6}, {3, 2, 2}, {5, 2, 8}},
			},
			4,
		},
		{
			"4",
			args{
				9,
				[][]int{{6, 2, 35129}, {3, 4, 99499}, {2, 7, 43547}, {8, 1, 78671}, {2, 1, 66308}, {9, 6, 33462}, {5, 1, 48249}, {2, 3, 44414}, {6, 7, 44602}, {1, 7, 14931}, {8, 9, 38171}, {4, 5, 30827}, {3, 9, 79166}, {4, 8, 93731}, {5, 9, 64068}, {7, 5, 17741}, {6, 3, 76017}, {9, 4, 72244}},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRestrictedPaths(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countRestrictedPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

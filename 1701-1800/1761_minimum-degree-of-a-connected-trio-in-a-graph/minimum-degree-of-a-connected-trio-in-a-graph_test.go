package main

import "testing"

func Test_minTrioDegree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{6, [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}}, 3},
		{"2", args{7, [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTrioDegree(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("minTrioDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}

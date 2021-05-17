package main

import "testing"

func Test_maxSideLength(t *testing.T) {
	type args struct {
		mat       [][]int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, 4}, 2},
		{"2", args{[][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, 1}, 0},
		{"3", args{[][]int{{1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 6}, 3},
		{"4", args{[][]int{{18, 70}, {61, 1}, {25, 85}, {14, 40}, {11, 96}, {97, 96}, {63, 45}}, 40184}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSideLength(tt.args.mat, tt.args.threshold); got != tt.want {
				t.Errorf("maxSideLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

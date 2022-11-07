package main

import "testing"

func Test_largestOverlap(t *testing.T) {
	type args struct {
		img1 [][]int
		img2 [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}}, 3},
		{"2", args{[][]int{{1}}, [][]int{{1}}}, 1},
		{"3", args{[][]int{{0}}, [][]int{{0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestOverlap(tt.args.img1, tt.args.img2); got != tt.want {
				t.Errorf("largestOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

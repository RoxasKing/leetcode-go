package main

import "testing"

func Test_largestSubmatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}}, 4},
		{"2", args{[][]int{{1, 0, 1, 0, 1}}}, 3},
		{"3", args{[][]int{{1, 1, 0}, {1, 0, 1}}}, 2},
		{"4", args{[][]int{{0, 0}, {0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSubmatrix(tt.args.matrix); got != tt.want {
				t.Errorf("largestSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

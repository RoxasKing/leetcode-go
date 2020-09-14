package main

import (
	"testing"
)

func Test_longestIncreasingPath(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}}, 4},
		{"2", args{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestIncreasingPath2(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}}, 4},
		{"2", args{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath2(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath2() = %v, want %v", got, tt.want)
			}
		})
	}
}

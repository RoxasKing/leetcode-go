package main

import (
	"testing"
)

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{2, [][]int{{1, 0}}}, true},
		{"2", args{2, [][]int{{1, 0}, {0, 1}}}, false},
		{"3", args{3, [][]int{{1, 0}, {1, 2}, {0, 1}}}, false},
		{"4", args{3, [][]int{{0, 2}, {1, 2}, {2, 0}}}, false},
		{"5", args{2, [][]int{{0, 1}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}

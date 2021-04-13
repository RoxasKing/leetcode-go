package main

import "testing"

func Test_processTasks(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}}, 4},
		{"2", args{[][]int{{2, 3, 1}, {5, 5, 1}, {5, 6, 2}}}, 3},
		{"3", args{[][]int{{1, 3, 2}, {2, 4, 2}, {3, 5, 2}, {4, 6, 2}, {5, 7, 2}, {6, 8, 2}, {7, 9, 2}}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processTasks(tt.args.tasks); got != tt.want {
				t.Errorf("processTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

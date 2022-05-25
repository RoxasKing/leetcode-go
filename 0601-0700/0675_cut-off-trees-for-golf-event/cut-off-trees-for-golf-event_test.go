package main

import "testing"

func Test_cutOffTree(t *testing.T) {
	type args struct {
		forest [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}}}, 6},
		{"2", args{[][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}}}, -1},
		{"3", args{[][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutOffTree(tt.args.forest); got != tt.want {
				t.Errorf("cutOffTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_maxCount(t *testing.T) {
	type args struct {
		m   int
		n   int
		ops [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 3, [][]int{{2, 2}, {3, 3}}}, 4},
		{"2", args{3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}}, 4},
		{"3", args{3, 3, [][]int{}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.args.m, tt.args.n, tt.args.ops); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

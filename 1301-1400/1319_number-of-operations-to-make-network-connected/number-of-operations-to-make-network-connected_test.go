package main

import "testing"

func Test_makeConnected(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{4, [][]int{{0, 1}, {0, 2}, {1, 2}}}, 1},
		{"2", args{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}}, 2},
		{"3", args{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}}, -1},
		{"4", args{5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeConnected(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("makeConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}

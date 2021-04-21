package main

import "testing"

func Test_maxNumEdgesToRemove(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}}, 2},
		{"2", args{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}}, 0},
		{"3", args{4, [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumEdgesToRemove(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("maxNumEdgesToRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}

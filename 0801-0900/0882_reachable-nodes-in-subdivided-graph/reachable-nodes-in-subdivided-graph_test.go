package main

import "testing"

func Test_reachableNodes(t *testing.T) {
	type args struct {
		edges    [][]int
		maxMoves int
		n        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}, 6, 3}, 13},
		{"2", args{[][]int{{0, 1, 4}, {1, 2, 6}, {0, 2, 8}, {1, 3, 1}}, 10, 4}, 23},
		{"3", args{[][]int{{1, 2, 4}, {1, 4, 5}, {1, 3, 1}, {2, 3, 4}, {3, 4, 5}}, 17, 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachableNodes(tt.args.edges, tt.args.maxMoves, tt.args.n); got != tt.want {
				t.Errorf("reachableNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

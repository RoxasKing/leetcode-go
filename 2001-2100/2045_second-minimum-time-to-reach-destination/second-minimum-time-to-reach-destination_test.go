package main

import "testing"

func Test_secondMinimum(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		time   int
		change int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5}, 13},
		{"2", args{2, [][]int{{1, 2}}, 3, 2}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondMinimum(tt.args.n, tt.args.edges, tt.args.time, tt.args.change); got != tt.want {
				t.Errorf("secondMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

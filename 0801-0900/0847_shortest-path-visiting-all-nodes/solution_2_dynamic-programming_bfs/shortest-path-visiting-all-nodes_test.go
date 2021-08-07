package main

import "testing"

func Test_shortestPathLength(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 2, 3}, {0}, {0}, {0}}}, 4},
		{"2", args{[][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}}}, 4},
		{"3", args{[][]int{{}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathLength(tt.args.graph); got != tt.want {
				t.Errorf("shortestPathLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

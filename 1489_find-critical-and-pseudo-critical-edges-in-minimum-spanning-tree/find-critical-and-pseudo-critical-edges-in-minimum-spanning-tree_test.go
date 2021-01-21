package main

import (
	"reflect"
	"testing"
)

func Test_findCriticalAndPseudoCriticalEdges(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				5,
				[][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}},
			},
			[][]int{{0, 1}, {2, 3, 4, 5}},
		},
		{
			"2",
			args{
				4,
				[][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}},
			},
			[][]int{nil, {0, 1, 2, 3}},
		},
		{
			"3",
			args{
				6,
				[][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 1}, {2, 3, 4}, {3, 4, 2}, {3, 5, 2}, {4, 5, 2}},
			},
			[][]int{{3}, {0, 1, 2, 4, 5, 6}},
		},
		{
			"4",
			args{
				6,
				[][]int{{0, 1, 1}, {0, 3, 1}, {0, 2, 1}, {1, 2, 1}, {1, 3, 1}, {2, 3, 1}},
			},
			[][]int{nil, {0, 1, 2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCriticalAndPseudoCriticalEdges(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCriticalAndPseudoCriticalEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}

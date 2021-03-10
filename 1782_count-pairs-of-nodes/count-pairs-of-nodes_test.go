package main

import (
	"reflect"
	"testing"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		n       int
		edges   [][]int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1",
			args{
				4,
				[][]int{{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1}},
				[]int{2, 3},
			},
			[]int{6, 5},
		},
		{
			"2",
			args{
				5,
				[][]int{{1, 5}, {1, 5}, {3, 4}, {2, 5}, {1, 3}, {5, 1}, {2, 3}, {2, 5}},
				[]int{1, 2, 3, 4, 5},
			},
			[]int{10, 10, 9, 8, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.n, tt.args.edges, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_sortItems(t *testing.T) {
	type args struct {
		n           int
		m           int
		group       []int
		beforeItems [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				8,
				2,
				[]int{-1, -1, 1, 0, 0, 1, 0, -1},
				[][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}},
			},
			[]int{6, 3, 4, 5, 2, 0, 1, 7},
		},
		{
			"2",
			args{
				8,
				2,
				[]int{-1, -1, 1, 0, 0, 1, 0, -1},
				[][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}},
			},
			[]int{},
		},
		{
			"3",
			args{
				5,
				5,
				[]int{2, 0, -1, 3, 0},
				[][]int{{2, 1, 3}, {2, 4}, {}, {}, {}},
			},
			[]int{2, 4, 1, 3, 0},
		},
		{
			"4",
			args{
				5,
				5,
				[]int{0, 1, 2, 3, 4},
				[][]int{{1}, {0}, {2}, {3}, {4}},
			},
			[]int{},
		},
		{
			"5",
			args{
				4,
				1,
				[]int{-1, 0, 0, -1},
				[][]int{{}, {0}, {1, 3}, {2}},
			},
			[]int{},
		},
		{
			"6",
			args{
				10,
				4,
				[]int{2, 2, 2, 1, 0, 1, 3, 2, 0, 1},
				[][]int{{7, 6, 2, 5, 3}, {}, {}, {}, {7}, {}, {}, {}, {}, {}},
			},
			[]int{3, 5, 9, 6, 2, 7, 1, 0, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortItems(tt.args.n, tt.args.m, tt.args.group, tt.args.beforeItems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

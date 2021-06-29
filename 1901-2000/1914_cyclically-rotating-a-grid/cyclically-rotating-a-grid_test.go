package main

import (
	"reflect"
	"testing"
)

func Test_rotateGrid(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{40, 10}, {30, 20}}, 1}, [][]int{{10, 20}, {40, 30}}},
		{
			"2",
			args{
				[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
				2,
			},
			[][]int{{3, 4, 8, 12}, {2, 11, 10, 16}, {1, 7, 6, 15}, {5, 9, 13, 14}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateGrid(tt.args.grid, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

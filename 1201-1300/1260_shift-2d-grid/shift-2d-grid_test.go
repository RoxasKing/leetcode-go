package main

import (
	"reflect"
	"testing"
)

func Test_shiftGrid(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1}, [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}}},
		{"2", args{[][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4}, [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}}},
		{"3", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftGrid(tt.args.grid, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shiftGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

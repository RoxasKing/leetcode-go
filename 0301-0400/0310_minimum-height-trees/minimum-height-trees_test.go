package main

import (
	"reflect"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{4, [][]int{{1, 0}, {1, 2}, {1, 3}}}, []int{1}},
		{"2", args{6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}}, []int{3, 4}},
		{"3", args{1, [][]int{}}, []int{0}},
		{"4", args{2, [][]int{{0, 1}}}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinHeightTrees(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

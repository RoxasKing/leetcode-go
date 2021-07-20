package main

import (
	"reflect"
	"testing"
)

func Test_maxGeneticDifference(t *testing.T) {
	type args struct {
		parents []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{-1, 0, 1, 1}, [][]int{{0, 2}, {3, 2}, {2, 5}}}, []int{2, 3, 7}},
		{"2", args{[]int{3, 7, -1, 2, 0, 7, 0, 2}, [][]int{{4, 6}, {1, 15}, {0, 5}}}, []int{6, 14, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxGeneticDifference(tt.args.parents, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxGeneticDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

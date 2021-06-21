package main

import (
	"reflect"
	"testing"
)

func Test_minDifference(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}}}, []int{2, 1, 4, 1}},
		{"2", args{[]int{4, 5, 2, 2, 7, 10}, [][]int{{2, 3}, {0, 2}, {0, 5}, {3, 5}}}, []int{-1, 1, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifference(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

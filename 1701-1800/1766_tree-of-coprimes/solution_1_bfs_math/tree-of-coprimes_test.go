package main

import (
	"reflect"
	"testing"
)

func Test_getCoprimes(t *testing.T) {
	type args struct {
		nums  []int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{2, 3, 3, 2}, [][]int{{0, 1}, {1, 2}, {1, 3}}}, []int{-1, 0, 0, 1}},
		{
			"2",
			args{
				[]int{5, 6, 10, 2, 3, 6, 15},
				[][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			},
			[]int{-1, 0, -1, 0, 0, 0, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCoprimes(tt.args.nums, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

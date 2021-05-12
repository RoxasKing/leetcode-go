package main

import (
	"reflect"
	"testing"
)

func Test_xorQueries(t *testing.T) {
	type args struct {
		arr     []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}}, []int{2, 7, 14, 8}},
		{"2", args{[]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}}, []int{8, 0, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorQueries(tt.args.arr, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xorQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

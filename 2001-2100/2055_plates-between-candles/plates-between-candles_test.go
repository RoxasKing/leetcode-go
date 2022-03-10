package main

import (
	"reflect"
	"testing"
)

func Test_platesBetweenCandles(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"**|**|***|", [][]int{{2, 5}, {5, 9}}}, []int{2, 3}},
		{"2", args{"***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}}, []int{9, 0, 0, 0, 0}},
		{"3", args{"***", [][]int{{2, 2}}}, []int{0}},
		{"4", args{"*|*|||", [][]int{{0, 0}, {1, 3}}}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := platesBetweenCandles(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("platesBetweenCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_sumEvenAfterQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}}, []int{8, 6, 2, 4}},
		{"2", args{[]int{1}, [][]int{{4, 0}}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumEvenAfterQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumEvenAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

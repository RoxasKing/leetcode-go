package main

import (
	"reflect"
	"testing"
)

func Test_getBiggestThree(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{[][]int{{3, 4, 5, 1, 3}, {3, 3, 4, 2, 3}, {20, 30, 200, 40, 10}, {1, 5, 5, 4, 1}, {4, 3, 2, 2, 5}}},
			[]int{228, 216, 211},
		},
		{
			"2",
			args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			[]int{20, 9, 8},
		},
		{
			"3",
			args{[][]int{{7, 7, 7}}},
			[]int{7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBiggestThree(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBiggestThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

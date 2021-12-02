package main

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		firstList  [][]int
		secondList [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
				[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			"2",
			args{
				[][]int{{1, 3}, {5, 9}},
				[][]int{},
			},
			[][]int{},
		},
		{
			"3",
			args{
				[][]int{},
				[][]int{{4, 8}, {10, 12}},
			},
			[][]int{},
		},
		{
			"4",
			args{
				[][]int{{1, 7}},
				[][]int{{3, 10}},
			},
			[][]int{{3, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalIntersection(tt.args.firstList, tt.args.secondList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_minimumAbsDifference(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{4, 2, 1, 3}}, [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{"2", args{[]int{1, 3, 6, 10, 15}}, [][]int{{1, 3}}},
		{"3", args{[]int{3, 8, -10, 23, 19, -4, -14, 27}}, [][]int{{-14, -10}, {19, 23}, {23, 27}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAbsDifference(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumAbsDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

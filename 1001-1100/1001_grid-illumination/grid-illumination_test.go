package main

import (
	"reflect"
	"testing"
)

func Test_gridIllumination(t *testing.T) {
	type args struct {
		N       int
		lamps   [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 0}}}, []int{1, 0}},
		{"2", args{5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 1}}}, []int{1, 1}},
		{"3", args{5, [][]int{{0, 0}, {0, 4}}, [][]int{{0, 4}, {0, 1}, {1, 4}}}, []int{1, 1, 0}},
		{
			"4",
			args{
				6,
				[][]int{{2, 5}, {4, 2}, {0, 3}, {0, 5}, {1, 4}, {4, 2}, {3, 3}, {1, 0}},
				[][]int{{4, 3}, {3, 1}, {5, 3}, {0, 5}, {4, 4}, {3, 3}},
			},
			[]int{1, 0, 1, 1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gridIllumination(tt.args.N, tt.args.lamps, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gridIllumination() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{2, 5, 2, 1, 2}, 5}, [][]int{{1, 2, 2}, {5}}},
		{
			"2",
			args{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
			[][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.args.candidates, tt.args.target)
			sort.Slice(got, func(i, j int) bool {
				var index int
				for index < len(got[i]) && index < len(got[j]) && got[i][index] == got[j][index] {
					index++
				}
				if index == len(got[i]) || index == len(got[j]) {
					return len(got[i]) < len(got[j])
				}
				return got[i][index] < got[j][index]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				var index int
				for tt.want[i][index] == tt.want[j][index] {
					index++
				}
				return tt.want[i][index] < tt.want[j][index]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{[]int{1, 2, 3}},
			[][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.args.nums)
			sort.Slice(got, func(i, j int) bool {
				if len(got[i]) != len(got[j]) {
					return len(got[i]) < len(got[j])
				}
				for k := range got[i] {
					if got[i][k] == got[j][k] {
						continue
					}
					return got[i][k] < got[j][k]
				}
				return true
			})
			sort.Slice(tt.want, func(i, j int) bool {
				if len(tt.want[i]) != len(tt.want[j]) {
					return len(tt.want[i]) < len(tt.want[j])
				}
				for k := range tt.want[i] {
					if tt.want[i][k] == tt.want[j][k] {
						continue
					}
					return tt.want[i][k] < tt.want[j][k]
				}
				return true
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

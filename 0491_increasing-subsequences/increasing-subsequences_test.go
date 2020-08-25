package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
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
			args{[]int{4, 6, 7, 7}},
			[][]int{{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7}},
		},
		{
			"2",
			args{[]int{1, 2, 1, 1}},
			[][]int{{1, 2}, {1, 1}, {1, 1, 1}},
		},
		{
			"3",
			args{[]int{1, 2, 1, 1, 1}},
			[][]int{{1, 2}, {1, 1}, {1, 1, 1}, {1, 1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubsequences(tt.args.nums)
			sort.Slice(got, func(i, j int) bool {
				if len(got[i]) != len(got[j]) {
					return len(got[i]) < len(got[j])
				} else {
					var k int
					for k < len(got[i])-1 && got[i][k] == got[j][k] {
						k++
					}
					return got[i][k] < got[j][k]
				}
			})
			sort.Slice(tt.want, func(i, j int) bool {
				if len(tt.want[i]) != len(tt.want[j]) {
					return len(tt.want[i]) < len(tt.want[j])
				} else {
					var k int
					for k < len(tt.want[i])-1 && tt.want[i][k] == tt.want[j][k] {
						k++
					}
					return tt.want[i][k] < tt.want[j][k]
				}
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

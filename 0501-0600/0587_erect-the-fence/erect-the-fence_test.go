package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_outerTrees(t *testing.T) {
	type args struct {
		trees [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}}}, [][]int{{1, 1}, {2, 0}, {3, 3}, {2, 4}, {4, 2}}},
		{"2", args{[][]int{{1, 2}, {2, 2}, {4, 2}}}, [][]int{{4, 2}, {2, 2}, {1, 2}}},
		{"3", args{[][]int{{1, 2}, {2, 2}, {4, 2}, {5, 2}, {6, 2}, {7, 2}}}, [][]int{{1, 2}, {2, 2}, {4, 2}, {5, 2}, {6, 2}, {7, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := outerTrees(tt.args.trees)
			sort.Slice(got, func(i, j int) bool {
				return got[i][0] < got[j][0] || got[i][0] == got[j][0] && got[i][1] < got[j][1]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i][0] < tt.want[j][0] || tt.want[i][0] == tt.want[j][0] && tt.want[i][1] < tt.want[j][1]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outerTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

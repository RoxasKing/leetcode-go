package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 3}, {-2, 2}}, 1}, [][]int{{-2, 2}}},
		{"2", args{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2}, [][]int{{3, 3}, {-2, 4}}},
		{"3", args{[][]int{{-5, 4}, {6, -5}, {4, 6}}, 2}, [][]int{{-5, 4}, {4, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kClosest(tt.args.points, tt.args.K)
			sort.Slice(got, func(i, j int) bool {
				return got[i][0] < got[j][0] || got[i][0] == got[j][0] && got[i][1] < got[j][1]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i][0] < tt.want[j][0] || tt.want[i][0] == tt.want[j][0] && tt.want[i][1] < tt.want[j][1]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

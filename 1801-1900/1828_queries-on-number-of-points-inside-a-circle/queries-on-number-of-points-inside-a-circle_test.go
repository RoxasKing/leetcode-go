package main

import (
	"reflect"
	"testing"
)

func Test_countPoints(t *testing.T) {
	type args struct {
		points  [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}}, []int{3, 2, 2}},
		{"2", args{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}}, []int{2, 3, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.args.points, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

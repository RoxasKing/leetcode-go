package main

import (
	"reflect"
	"testing"
)

func Test_restoreArray(t *testing.T) {
	type args struct {
		adjacentPairs [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{2, 1}, {3, 4}, {3, 2}}}, [][]int{{1, 2, 3, 4}, {4, 3, 2, 1}}},
		{"2", args{[][]int{{4, -2}, {1, 4}, {-3, 1}}}, [][]int{{-2, 4, 1, -3}, {-3, 1, 4, -2}}},
		{"3", args{[][]int{{100000, -100000}}}, [][]int{{100000, -100000}, {-100000, 100000}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := restoreArray(tt.args.adjacentPairs)
			if !reflect.DeepEqual(got, tt.want[0]) && !reflect.DeepEqual(got, tt.want[1]) {
				t.Errorf("restoreArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

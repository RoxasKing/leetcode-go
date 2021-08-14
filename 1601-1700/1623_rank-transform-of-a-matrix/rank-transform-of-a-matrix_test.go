package main

import (
	"reflect"
	"testing"
)

func Test_matrixRankTransform(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 2}, {3, 4}}}, [][]int{{1, 2}, {2, 3}}},
		{"2", args{[][]int{{7, 7}, {7, 7}}}, [][]int{{1, 1}, {1, 1}}},
		{"3", args{[][]int{{20, -21, 14}, {-19, 4, 19}, {22, -47, 24}, {-19, 4, 19}}}, [][]int{{4, 2, 3}, {1, 3, 4}, {5, 1, 6}, {1, 3, 4}}},
		{"4", args{[][]int{{7, 3, 6}, {1, 4, 5}, {9, 8, 2}}}, [][]int{{5, 1, 4}, {1, 2, 3}, {6, 3, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixRankTransform(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

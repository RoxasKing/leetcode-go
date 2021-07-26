package main

import (
	"reflect"
	"testing"
)

func Test_splitPainting(t *testing.T) {
	type args struct {
		segments [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int64
	}{
		{"1", args{[][]int{{1, 4, 5}, {4, 7, 7}, {1, 7, 9}}}, [][]int64{{1, 4, 14}, {4, 7, 16}}},
		{
			"2",
			args{[][]int{{1, 7, 9}, {6, 8, 15}, {8, 10, 7}}},
			[][]int64{{1, 6, 9}, {6, 7, 24}, {7, 8, 15}, {8, 10, 7}},
		},
		{"3", args{[][]int{{1, 4, 5}, {1, 4, 7}, {4, 7, 1}, {4, 7, 11}}}, [][]int64{{1, 4, 12}, {4, 7, 12}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitPainting(tt.args.segments); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitPainting() = %v, want %v", got, tt.want)
			}
		})
	}
}

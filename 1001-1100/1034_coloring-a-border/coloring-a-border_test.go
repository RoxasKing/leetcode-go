package main

import (
	"reflect"
	"testing"
)

func Test_colorBorder(t *testing.T) {
	type args struct {
		grid  [][]int
		row   int
		col   int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 1}, {1, 2}}, 0, 0, 3}, [][]int{{3, 3}, {3, 2}}},
		{"2", args{[][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3}, [][]int{{1, 3, 3}, {2, 3, 3}}},
		{"3", args{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1, 1, 2}, [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colorBorder(tt.args.grid, tt.args.row, tt.args.col, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("colorBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

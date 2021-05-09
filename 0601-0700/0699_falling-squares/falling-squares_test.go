package main

import (
	"reflect"
	"testing"
)

func Test_fallingSquares(t *testing.T) {
	type args struct {
		positions [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{1, 2}, {2, 3}, {6, 1}}}, []int{2, 5, 5}},
		{"2", args{[][]int{{100, 100}, {200, 100}}}, []int{100, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fallingSquares(tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fallingSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_findSolution(t *testing.T) {
	func1 := func(x, y int) int {
		return x + y
	}
	func2 := func(x, y int) int {
		return x * y
	}
	type args struct {
		customFunction func(int, int) int
		z              int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{func1, 5}, [][]int{{1, 4}, {2, 3}, {3, 2}, {4, 1}}},
		{"2", args{func2, 5}, [][]int{{1, 5}, {5, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSolution(tt.args.customFunction, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

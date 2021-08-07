package main

import (
	"reflect"
	"testing"
)

func Test_gardenNoAdj(t *testing.T) {
	type args struct {
		n     int
		paths [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{3, [][]int{{1, 2}, {2, 3}, {3, 1}}}, []int{1, 2, 3}},
		{"2", args{4, [][]int{{1, 2}, {3, 4}}}, []int{1, 2, 1, 2}},
		{"3", args{4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gardenNoAdj(tt.args.n, tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gardenNoAdj() = %v, want %v", got, tt.want)
			}
		})
	}
}

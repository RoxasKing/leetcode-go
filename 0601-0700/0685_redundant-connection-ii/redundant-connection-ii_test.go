package main

import (
	"reflect"
	"testing"
)

func Test_findRedundantDirectedConnection(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{1, 2}, {1, 3}, {2, 3}}}, []int{2, 3}},
		{"2", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}}, []int{4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRedundantDirectedConnection(tt.args.edges)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantDirectedConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

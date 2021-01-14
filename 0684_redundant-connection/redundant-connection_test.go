package main

import (
	"reflect"
	"testing"
)

func Test_findRedundantConnection(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 2}, {1, 3}, {2, 3}}}, [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{"2", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}}, [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}}},
		{"3", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {2, 4}, {1, 5}}}, [][]int{{2, 3}, {3, 4}, {2, 4}}},
		{"4", args{[][]int{{1, 2}, {2, 3}, {1, 3}, {3, 4}, {1, 5}}}, [][]int{{1, 2}, {2, 3}, {1, 3}}},
		{"5", args{[][]int{{1, 2}, {2, 3}, {2, 4}, {3, 4}, {1, 5}}}, [][]int{{2, 3}, {3, 4}, {2, 4}}},
		{"6", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}, {1, 5}}}, [][]int{{1, 2}, {2, 3}, {1, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRedundantConnection(tt.args.edges)
			contain := false
		LOOP:
			for _, want := range tt.want {
				if reflect.DeepEqual(got, want) {
					contain = true
					break LOOP
				}
			}
			if !contain {
				t.Errorf("findRedundantConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

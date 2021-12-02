package main

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 2}, {3}, {3}, {}}}, [][]int{{0, 1, 3}, {0, 2, 3}}},
		{"2", args{[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}}, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}},
		{"3", args{[][]int{{1}, {}}}, [][]int{{0, 1}}},
		{"4", args{[][]int{{1, 2, 3}, {2}, {3}, {}}}, [][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}}},
		{"5", args{[][]int{{1, 3}, {2}, {3}, {}}}, [][]int{{0, 1, 2, 3}, {0, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_construct2DArray(t *testing.T) {
	type args struct {
		original []int
		m        int
		n        int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{1, 2, 3, 4}, 2, 2}, [][]int{{1, 2}, {3, 4}}},
		{"2", args{[]int{1, 2, 3}, 1, 3}, [][]int{{1, 2, 3}}},
		{"3", args{[]int{1, 2}, 1, 1}, nil},
		{"4", args{[]int{3}, 1, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct2DArray(tt.args.original, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

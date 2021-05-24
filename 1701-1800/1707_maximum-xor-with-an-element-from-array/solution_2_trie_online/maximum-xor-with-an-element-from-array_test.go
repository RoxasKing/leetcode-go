package main

import (
	"reflect"
	"testing"
)

func Test_maximizeXor(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{0, 1, 2, 3, 4}, [][]int{{3, 1}, {1, 3}, {5, 6}}}, []int{3, 3, 7}},
		{"2", args{[]int{5, 2, 4, 6, 6, 3}, [][]int{{12, 4}, {8, 1}, {6, 3}}}, []int{15, -1, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeXor(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximizeXor() = %v, want %v", got, tt.want)
			}
		})
	}
}

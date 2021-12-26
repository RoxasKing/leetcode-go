package main

import (
	"reflect"
	"testing"
)

func Test_loudAndRich(t *testing.T) {
	type args struct {
		richer [][]int
		quiet  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}}, []int{5, 5, 2, 5, 4, 5, 6, 7}},
		{"2", args{[][]int{}, []int{0}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loudAndRich(tt.args.richer, tt.args.quiet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loudAndRich() = %v, want %v", got, tt.want)
			}
		})
	}
}

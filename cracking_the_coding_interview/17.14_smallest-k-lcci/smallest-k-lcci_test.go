package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_smallestK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 3, 5, 7, 2, 4, 6, 8}, 4}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestK(tt.args.arr, tt.args.k)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestK() = %v, want %v", got, tt.want)
			}
		})
	}
}

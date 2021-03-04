package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 2, 1}, 2}, []int{1, 2}},
		{"2", args{[]int{0, 1, 2, 1}, 1}, []int{0}},
		{"3", args{[]int{0, 0, 2, 3, 2, 1, 1, 2, 0, 4}, 10}, []int{0, 0, 0, 1, 1, 2, 2, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLeastNumbers(tt.args.arr, tt.args.k)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

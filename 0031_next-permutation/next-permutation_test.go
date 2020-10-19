package main

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3}}, []int{1, 3, 2}},
		{"2", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"3", args{[]int{1, 1, 5}}, []int{1, 5, 1}},
		{"4", args{[]int{1, 3, 2, 4}}, []int{1, 3, 4, 2}},
		{"5", args{[]int{1, 3, 4, 2}}, []int{1, 4, 2, 3}},
		{"6", args{[]int{1, 4, 2, 3}}, []int{1, 4, 3, 2}},
		{"7", args{[]int{1, 4, 3, 2}}, []int{2, 1, 3, 4}},
		{"8", args{[]int{1, 1, 2, 2}}, []int{1, 2, 1, 2}},
		{"9", args{[]int{1, 1, 1, 1}}, []int{1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

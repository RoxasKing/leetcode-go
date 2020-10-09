package main

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{2, 0, 2, 1, 1, 0}}, []int{0, 0, 1, 1, 2, 2}},
		{"2", args{[]int{2, 0, 1}}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

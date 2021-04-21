package main

import (
	"reflect"
	"testing"
)

func Test_getMaximumXor(t *testing.T) {
	type args struct {
		nums       []int
		maximumBit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{0, 1, 1, 3}, 2}, []int{0, 3, 2, 3}},
		{"2", args{[]int{2, 3, 4, 7}, 3}, []int{5, 2, 6, 5}},
		{"3", args{[]int{0, 1, 2, 2, 5, 7}, 3}, []int{4, 3, 6, 4, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumXor(tt.args.nums, tt.args.maximumBit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMaximumXor() = %v, want %v", got, tt.want)
			}
		})
	}
}

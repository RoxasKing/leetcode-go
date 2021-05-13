package main

import (
	"reflect"
	"testing"
)

func Test_threeEqualParts(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 0, 1, 0, 1}}, []int{0, 3}},
		{"2", args{[]int{1, 1, 0, 1, 1}}, []int{-1, -1}},
		{"3", args{[]int{1, 1, 0, 0, 1}}, []int{0, 2}},
		{"4", args{[]int{1, 1, 1, 0}}, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeEqualParts(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeEqualParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

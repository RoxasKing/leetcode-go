package main

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"2", args{[]int{}}, []int{}},
		{"3", args{[]int{0, 0}}, []int{0, 0}},
		{"4", args{[]int{1, 0}}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

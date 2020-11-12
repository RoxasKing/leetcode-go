package main

import "testing"

func Test_maxDotProduct(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 1, -2, 5}, []int{3, 0, -6}}, 18},
		{"2", args{[]int{3, -2}, []int{2, -6, 7}}, 21},
		{"3", args{[]int{-1, -1}, []int{1, 1}}, -1},
		{"4", args{[]int{2, 1, -2, 5, -3}, []int{3, 0, -6, -9}}, 45},
		{"5", args{[]int{2, 1, -2, 5, -3}, []int{3, 0, -6, -9, -10}}, 54},
		{"6", args{[]int{-1, -1}, []int{1, -1, 1, -1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDotProduct(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxDotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

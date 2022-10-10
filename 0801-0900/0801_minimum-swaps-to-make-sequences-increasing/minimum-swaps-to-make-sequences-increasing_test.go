package main

import "testing"

func Test_minSwap(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 5, 4}, []int{1, 2, 7, 3}}, 1},
		{"2", args{[]int{0, 3, 5, 8, 9}, []int{2, 1, 4, 6, 9}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwap(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}

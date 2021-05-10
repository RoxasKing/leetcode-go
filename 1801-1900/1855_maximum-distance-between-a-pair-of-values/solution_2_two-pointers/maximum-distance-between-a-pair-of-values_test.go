package main

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}}, 2},
		{"2", args{[]int{2, 2, 2}, []int{10, 10, 1}}, 1},
		{"3", args{[]int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}}, 2},
		{"4", args{[]int{5, 4}, []int{3, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

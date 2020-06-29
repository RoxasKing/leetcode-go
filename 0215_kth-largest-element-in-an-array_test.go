package leetcode

import (
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"2", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4}, 4},
		{"3", args{[]int{99, 99}, 1}, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKthLargest2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"2", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargestInHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

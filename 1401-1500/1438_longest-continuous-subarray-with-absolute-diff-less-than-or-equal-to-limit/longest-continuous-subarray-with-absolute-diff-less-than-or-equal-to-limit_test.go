package main

import (
	"testing"
)

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{8, 2, 4, 7}, 4}, 2},
		{"2", args{[]int{10, 1, 2, 4, 7, 2}, 5}, 4},
		{"3", args{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0}, 3},
		{"4", args{[]int{8}, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubarray2(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{8, 2, 4, 7}, 4}, 2},
		{"2", args{[]int{10, 1, 2, 4, 7, 2}, 5}, 4},
		{"3", args{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0}, 3},
		{"4", args{[]int{8}, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray2(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 4, 2, 3}, 5}, 2},
		{"2", args{[]int{5, 6, 7, 8, 9}, 4}, -1},
		{"3", args{[]int{3, 2, 20, 1, 1, 3}, 10}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minOperations2(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 4, 2, 3}, 5}, 2},
		{"2", args{[]int{5, 6, 7, 8, 9}, 4}, -1},
		{"3", args{[]int{3, 2, 20, 1, 1, 3}, 10}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations2(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minOperations2() = %v, want %v", got, tt.want)
			}
		})
	}
}

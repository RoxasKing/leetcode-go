package main

import (
	"testing"
)

func Test_canPartitionKSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{4, 3, 2, 3, 5, 2, 1}, 4}, true},
		{"2", args{[]int{1, 2, 3, 4}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionKSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("canPartitionKSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartitionKSubsets2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{4, 3, 2, 3, 5, 2, 1}, 4}, true},
		{"2", args{[]int{1, 2, 3, 4}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionKSubsets2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("canPartitionKSubsets2() = %v, want %v", got, tt.want)
			}
		})
	}
}

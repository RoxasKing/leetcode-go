package leetcode

import (
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 1, 2, 3}}, 5},
		{"2", args{[]int{2, 1}}, 3},
		{"3", args{[]int{1, 2, 0}}, 3},
		{"4", args{[]int{3, 4, -1, 1}}, 2},
		{"5", args{[]int{7, 8, 9, 11, 12}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstMissingPositive2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 1, 2, 3}}, 5},
		{"2", args{[]int{2, 1}}, 3},
		{"3", args{[]int{1, 2, 0}}, 3},
		{"4", args{[]int{3, 4, -1, 1}}, 2},
		{"5", args{[]int{7, 8, 9, 11, 12}}, 1},
		{"6", args{[]int{-4, 24, 32, 25, 16, -8, 3, -5, -6, 30, 3, 3, 29, -5, 6, -3, 1, 29, -2, 4, 4, 7, 14, 20, 5, 0, 25, 2, 13, 26, -9, 7, 6, 33}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive2(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive2() = %v, want %v", got, tt.want)
			}
		})
	}
}

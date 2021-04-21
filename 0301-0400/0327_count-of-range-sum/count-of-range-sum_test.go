package main

import (
	"testing"
)

func Test_countRangeSum(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{-2, 5, -1}, -2, 2}, 3},
		{"2", args{[]int{-2, 5, -1, -1}, -2, 2}, 6},
		{"3", args{[]int{-1, 1}, 0, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRangeSum(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countRangeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countRangeSum2(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{-2, 5, -1}, -2, 2}, 3},
		{"2", args{[]int{-2, 5, -1, -1}, -2, 2}, 6},
		{"3", args{[]int{-1, 1}, 0, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRangeSum2(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countRangeSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

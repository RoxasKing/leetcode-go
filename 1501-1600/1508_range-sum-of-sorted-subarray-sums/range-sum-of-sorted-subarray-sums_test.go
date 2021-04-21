package main

import (
	"testing"
)

func Test_rangeSum(t *testing.T) {
	type args struct {
		nums  []int
		n     int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4}, 4, 1, 5}, 13},
		{"2", args{[]int{1, 2, 3, 4}, 4, 3, 4}, 6},
		{"3", args{[]int{1, 2, 3, 4}, 4, 1, 10}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSum(tt.args.nums, tt.args.n, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rangeSum2(t *testing.T) {
	type args struct {
		nums  []int
		n     int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4}, 4, 1, 5}, 13},
		{"2", args{[]int{1, 2, 3, 4}, 4, 3, 4}, 6},
		{"3", args{[]int{1, 2, 3, 4}, 4, 1, 10}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSum2(tt.args.nums, tt.args.n, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rangeSum3(t *testing.T) {
	type args struct {
		nums  []int
		n     int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4}, 4, 1, 5}, 13},
		{"2", args{[]int{1, 2, 3, 4}, 4, 3, 4}, 6},
		{"3", args{[]int{1, 2, 3, 4}, 4, 1, 10}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSum3(tt.args.nums, tt.args.n, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}

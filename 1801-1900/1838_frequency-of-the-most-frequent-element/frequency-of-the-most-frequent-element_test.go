package main

import (
	"testing"
)

func Test_maxFrequency(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 4}, 5}, 3},
		{"2", args{[]int{1, 4, 8, 13}, 5}, 2},
		{"3", args{[]int{3, 9, 6}, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxFrequency2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 4}, 5}, 3},
		{"2", args{[]int{1, 4, 8, 13}, 5}, 2},
		{"3", args{[]int{3, 9, 6}, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxFrequency2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{10, 9, 2, 5, 3, 4}}, 3},
		{"", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{0}}, 1},
		{"", args{[]int{10, 9, 2, 5, 3, 4}}, 3},
		{"", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS2(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS2() = %v, want %v", got, tt.want)
			}
		})
	}
}

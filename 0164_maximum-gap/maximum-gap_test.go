package main

import (
	"testing"
)

func Test_maximumGap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 6, 9, 1}}, 3},
		{"2", args{[]int{10}}, 0},
		{"3", args{[]int{1, 10000000}}, 9999999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGap(tt.args.nums); got != tt.want {
				t.Errorf("maximumGap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumGap2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 6, 9, 1}}, 3},
		{"2", args{[]int{10}}, 0},
		{"3", args{[]int{1, 10000000}}, 9999999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGap2(tt.args.nums); got != tt.want {
				t.Errorf("maximumGap2() = %v, want %v", got, tt.want)
			}
		})
	}
}

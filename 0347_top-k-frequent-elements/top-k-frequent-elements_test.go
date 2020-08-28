package main

import (
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
		{"2", args{[]int{1}, 1}, []int{1}},
		{"3", args{[]int{3, 0, 1, 0}, 1}, []int{0}},
		{"4", args{[]int{4, 1, -1, 2, -1, 2, 3}, 2}, []int{-1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.args.nums, tt.args.k)
			mark := make(map[int]bool)
			for _, num := range got {
				mark[num] = true
			}
			for _, num := range tt.want {
				if !mark[num] {
					t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_topKFrequent2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
		{"2", args{[]int{1}, 1}, []int{1}},
		{"3", args{[]int{3, 0, 1, 0}, 1}, []int{0}},
		{"4", args{[]int{4, 1, -1, 2, -1, 2, 3}, 2}, []int{-1, 2}},
		{"5", args{[]int{5, 3, 1, 1, 1, 3, 73, 1}, 2}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent2(tt.args.nums, tt.args.k)
			mark := make(map[int]bool)
			for _, num := range got {
				mark[num] = true
			}
			for _, num := range tt.want {
				if !mark[num] {
					t.Errorf("topKFrequent2() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

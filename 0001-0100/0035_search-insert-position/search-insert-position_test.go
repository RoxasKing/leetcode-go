package main

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{}, 1}, 0},
		{"2", args{[]int{1}, 0}, 0},
		{"3", args{[]int{1, 3, 5, 6}, 1}, 0},
		{"4", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"5", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"6", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"7", args{[]int{1, 3, 5, 6}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

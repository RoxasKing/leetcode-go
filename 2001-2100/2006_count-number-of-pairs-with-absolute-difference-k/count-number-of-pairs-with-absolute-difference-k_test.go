package main

import "testing"

func Test_countKDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 2, 1}, 1}, 4},
		{"2", args{[]int{1, 3}, 3}, 0},
		{"3", args{[]int{3, 2, 1, 5, 4}, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countKDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_minAbsDifference(t *testing.T) {
	type args struct {
		nums []int
		goal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{5, -7, 3, 5}, 6}, 0},
		{"2", args{[]int{7, -9, 15, -2}, -5}, 1},
		{"3", args{[]int{1, 2, 3}, -7}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAbsDifference(tt.args.nums, tt.args.goal); got != tt.want {
				t.Errorf("minAbsDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

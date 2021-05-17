package main

import "testing"

func Test_minSubarray(t *testing.T) {
	type args struct {
		nums []int
		p    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 1, 4, 2}, 6}, 1},
		{"2", args{[]int{6, 3, 5, 2}, 9}, 2},
		{"3", args{[]int{1, 2, 3}, 3}, 0},
		{"4", args{[]int{1, 2, 3}, 7}, -1},
		{"5", args{[]int{1000000000, 1000000000, 1000000000}, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubarray(tt.args.nums, tt.args.p); got != tt.want {
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

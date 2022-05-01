package main

import "testing"

func Test_smallestRangeI(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1}, 0}, 0},
		{"2", args{[]int{0, 10}, 2}, 6},
		{"3", args{[]int{1, 3, 6}, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRangeI(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestRangeI() = %v, want %v", got, tt.want)
			}
		})
	}
}

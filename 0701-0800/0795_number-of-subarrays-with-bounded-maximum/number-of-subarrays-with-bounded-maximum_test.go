package main

import "testing"

func Test_numSubarrayBoundedMax(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 1, 4, 3}, 2, 3}, 3},
		{"2", args{[]int{2, 9, 2, 5, 6}, 2, 8}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayBoundedMax(tt.args.nums, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("numSubarrayBoundedMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

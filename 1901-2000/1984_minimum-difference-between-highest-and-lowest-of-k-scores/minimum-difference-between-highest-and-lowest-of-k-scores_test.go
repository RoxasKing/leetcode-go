package main

import "testing"

func Test_minimumDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{90}, 1}, 0},
		{"2", args{[]int{9, 4, 1, 7}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

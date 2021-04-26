package main

import "testing"

func Test_smallestDistancePair(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 1}, 1}, 0},
		{"2", args{[]int{1, 3, 1, 5, 6, 10, 32, 43, 11}, 6}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestDistancePair(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestDistancePair() = %v, want %v", got, tt.want)
			}
		})
	}
}

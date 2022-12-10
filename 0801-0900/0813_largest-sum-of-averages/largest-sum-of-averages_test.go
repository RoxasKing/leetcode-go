package main

import "testing"

func Test_largestSumOfAverages(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{[]int{9, 1, 2, 3, 9}, 3}, 20.000000},
		{"2", args{[]int{1, 2, 3, 4, 5, 6, 7}, 4}, 20.500000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumOfAverages(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumOfAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}

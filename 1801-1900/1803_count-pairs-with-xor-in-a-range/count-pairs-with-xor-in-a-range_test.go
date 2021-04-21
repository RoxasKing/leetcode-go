package main

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 4, 2, 7}, 2, 6}, 6},
		{"2", args{[]int{9, 8, 4, 2, 1}, 5, 14}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.nums, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

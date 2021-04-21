package main

import "testing"

func Test_maximumScore(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 4, 3, 7, 4, 5}, 3}, 15},
		{"2", args{[]int{5, 5, 4, 5, 4, 1, 1, 1}, 0}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

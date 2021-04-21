package main

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5, 6, 1}, 3}, 12},
		{"2", args{[]int{2, 2, 2}, 2}, 4},
		{"3", args{[]int{9, 7, 7, 9, 7, 7, 9}, 7}, 55},
		{"4", args{[]int{1, 1000, 1}, 1}, 1},
		{"5", args{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3}, 202},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

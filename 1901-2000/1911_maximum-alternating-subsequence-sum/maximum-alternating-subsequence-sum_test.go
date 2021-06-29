package main

import "testing"

func Test_maxAlternatingSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{[]int{4, 2, 5, 3}}, 7},
		{"2", args{[]int{5, 6, 7, 8}}, 8},
		{"3", args{[]int{6, 2, 1, 2, 4, 5}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAlternatingSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAlternatingSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

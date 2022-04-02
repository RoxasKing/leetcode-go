package main

import "testing"

func Test_splitArray(t *testing.T) {
	type args struct {
		nums []int
		m    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{7, 2, 5, 10, 8}, 2}, 18},
		{"2", args{[]int{1, 2, 3, 4, 5}, 2}, 9},
		{"3", args{[]int{1, 4, 4}, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArray(tt.args.nums, tt.args.m); got != tt.want {
				t.Errorf("splitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_subArrayRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{[]int{1, 2, 3}}, 4},
		{"2", args{[]int{1, 3, 3}}, 4},
		{"3", args{[]int{4, -2, -3, 4, 1}}, 59},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subArrayRanges(tt.args.nums); got != tt.want {
				t.Errorf("subArrayRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

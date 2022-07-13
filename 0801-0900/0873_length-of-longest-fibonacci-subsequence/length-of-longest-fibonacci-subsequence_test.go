package main

import "testing"

func Test_lenLongestFibSubseq(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, 5},
		{"2", args{[]int{1, 3, 7, 11, 12, 14, 18}}, 3},
		{"3", args{[]int{2, 4, 5, 6, 7, 8, 11, 13, 14, 15, 21, 22, 34}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenLongestFibSubseq(tt.args.arr); got != tt.want {
				t.Errorf("lenLongestFibSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}

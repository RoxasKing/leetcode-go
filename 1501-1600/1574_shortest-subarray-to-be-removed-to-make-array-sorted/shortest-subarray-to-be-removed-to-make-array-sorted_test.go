package main

import "testing"

func Test_findLengthOfShortestSubarray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 10, 4, 2, 3, 5}}, 3},
		{"2", args{[]int{5, 4, 3, 2, 1}}, 4},
		{"3", args{[]int{1, 2, 3}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfShortestSubarray(tt.args.arr); got != tt.want {
				t.Errorf("findLengthOfShortestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

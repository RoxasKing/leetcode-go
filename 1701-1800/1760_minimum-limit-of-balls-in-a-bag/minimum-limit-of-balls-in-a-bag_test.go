package main

import "testing"

func Test_minimumSize(t *testing.T) {
	type args struct {
		nums          []int
		maxOperations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{9}, 2}, 3},
		{"2", args{[]int{2, 4, 8, 2}, 4}, 2},
		{"3", args{[]int{7, 17}, 2}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSize(tt.args.nums, tt.args.maxOperations); got != tt.want {
				t.Errorf("minimumSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

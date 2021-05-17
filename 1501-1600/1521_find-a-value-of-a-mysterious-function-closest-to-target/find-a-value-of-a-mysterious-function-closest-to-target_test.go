package main

import "testing"

func Test_closestToTarget(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{9, 12, 3, 7, 15}, 5}, 2},
		{"2", args{[]int{1000000, 1000000, 1000000}, 1}, 999999},
		{"3", args{[]int{1, 2, 4, 8, 16}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestToTarget(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("closestToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

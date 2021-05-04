package main

import "testing"

func Test_getMinDistance(t *testing.T) {
	type args struct {
		nums   []int
		target int
		start  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5}, 5, 3}, 1},
		{"2", args{[]int{1}, 1, 0}, 0},
		{"3", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinDistance(tt.args.nums, tt.args.target, tt.args.start); got != tt.want {
				t.Errorf("getMinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

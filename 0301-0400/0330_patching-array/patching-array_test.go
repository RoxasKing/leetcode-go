package main

import "testing"

func Test_minPatches(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3}, 6}, 1},
		{"2", args{[]int{1, 5, 10}, 20}, 2},
		{"3", args{[]int{1, 2, 2}, 5}, 0},
		{"4", args{[]int{1, 2, 31, 33}, 2147483647}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPatches(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("minPatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

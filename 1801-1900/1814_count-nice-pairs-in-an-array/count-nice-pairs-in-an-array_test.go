package main

import "testing"

func Test_countNicePairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{42, 11, 1, 97}}, 2},
		{"2", args{[]int{13, 10, 35, 24, 76}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNicePairs(tt.args.nums); got != tt.want {
				t.Errorf("countNicePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

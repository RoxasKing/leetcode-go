package main

import "testing"

func Test_minimumAverageDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 5, 3, 9, 5, 3}}, 3},
		{"2", args{[]int{0}}, 0},
		{"3", args{[]int{0, 1, 0, 1, 0, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAverageDifference(tt.args.nums); got != tt.want {
				t.Errorf("minimumAverageDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

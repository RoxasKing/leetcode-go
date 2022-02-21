package main

import "testing"

func Test_minimumDeviation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 5}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeviation(tt.args.nums); got != tt.want {
				t.Errorf("minimumDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}

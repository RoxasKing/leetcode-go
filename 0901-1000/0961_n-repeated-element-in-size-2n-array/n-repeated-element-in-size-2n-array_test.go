package main

import "testing"

func Test_repeatedNTimes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 3}}, 3},
		{"2", args{[]int{2, 1, 2, 5, 3, 2}}, 2},
		{"3", args{[]int{5, 1, 5, 2, 5, 3, 5, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedNTimes(tt.args.nums); got != tt.want {
				t.Errorf("repeatedNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_canBeIncreasing(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 10, 5, 7}}, true},
		{"2", args{[]int{2, 3, 1, 2}}, false},
		{"3", args{[]int{1, 1, 1}}, false},
		{"4", args{[]int{1, 2, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeIncreasing(tt.args.nums); got != tt.want {
				t.Errorf("canBeIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}

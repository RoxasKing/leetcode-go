package main

import "testing"

func Test_sumOfUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 2}}, 4},
		{"2", args{[]int{1, 1, 1, 1, 1}}, 0},
		{"3", args{[]int{1, 2, 3, 4, 5}}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfUnique(tt.args.nums); got != tt.want {
				t.Errorf("sumOfUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

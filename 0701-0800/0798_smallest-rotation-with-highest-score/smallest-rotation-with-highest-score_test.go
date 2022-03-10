package main

import "testing"

func Test_bestRotation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 3, 1, 4, 0}}, 3},
		{"2", args{[]int{1, 3, 0, 2, 4}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestRotation(tt.args.nums); got != tt.want {
				t.Errorf("bestRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_arraySign(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{-1, -2, -3, -4, 3, 2, 1}}, 1},
		{"2", args{[]int{1, 5, 0, 2, -3}}, 0},
		{"3", args{[]int{-1, 1, -1, 1, -1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arraySign(tt.args.nums); got != tt.want {
				t.Errorf("arraySign() = %v, want %v", got, tt.want)
			}
		})
	}
}

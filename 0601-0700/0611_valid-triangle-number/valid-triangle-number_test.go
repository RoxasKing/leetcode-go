package main

import "testing"

func Test_triangleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 2, 3, 4}}, 3},
		{"2", args{[]int{4, 2, 3, 4}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleNumber(tt.args.nums); got != tt.want {
				t.Errorf("triangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

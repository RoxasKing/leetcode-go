package leetcode

import (
	"testing"
)

func Test_findLength(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLength2(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength2(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("findLength2() = %v, want %v", got, tt.want)
			}
		})
	}
}

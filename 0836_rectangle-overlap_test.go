package main

import (
	"testing"
)

func Test_isRectangleOverlap(t *testing.T) {
	type args struct {
		rec1 []int
		rec2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{
				[]int{0, 0, 2, 2},
				[]int{1, 1, 3, 3},
			},
			true,
		},
		{
			"",
			args{
				[]int{0, 0, 1, 1},
				[]int{1, 0, 2, 1},
			},
			false,
		},
		{
			"",
			args{
				[]int{7, 8, 13, 15},
				[]int{10, 8, 12, 20},
			},
			true,
		},
		{
			"",
			args{
				[]int{2, 7, 6, 20},
				[]int{3, 8, 6, 20},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleOverlap(tt.args.rec1, tt.args.rec2); got != tt.want {
				t.Errorf("isRectangleOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRectangleOverlap2(t *testing.T) {
	type args struct {
		rec1 []int
		rec2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{
				[]int{0, 0, 2, 2},
				[]int{1, 1, 3, 3},
			},
			true,
		},
		{
			"",
			args{
				[]int{0, 0, 1, 1},
				[]int{1, 0, 2, 1},
			},
			false,
		},
		{
			"",
			args{
				[]int{7, 8, 13, 15},
				[]int{10, 8, 12, 20},
			},
			true,
		},
		{
			"",
			args{
				[]int{2, 7, 6, 20},
				[]int{3, 8, 6, 20},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleOverlap2(tt.args.rec1, tt.args.rec2); got != tt.want {
				t.Errorf("isRectangleOverlap2() = %v, want %v", got, tt.want)
			}
		})
	}
}

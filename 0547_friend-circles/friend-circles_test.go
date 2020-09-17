package main

import (
	"testing"
)

func Test_findCircleNum(t *testing.T) {
	type args struct {
		M [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[][]int{
					{1, 1, 0},
					{1, 1, 0},
					{0, 0, 1},
				},
			},
			2,
		},
		{
			"2",
			args{
				[][]int{
					{1, 1, 0},
					{1, 1, 1},
					{0, 1, 1},
				},
			},
			1,
		},
		{
			"3",
			args{
				[][]int{
					{1, 0, 0, 1},
					{0, 1, 1, 0},
					{0, 1, 1, 1},
					{1, 0, 1, 1},
				},
			},
			1,
		},
		{
			"4",
			args{
				[][]int{
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
				},
			},
			1,
		},
		{
			"5",
			args{
				[][]int{
					{1, 1, 1, 0},
					{1, 1, 0, 0},
					{1, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.args.M); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCircleNum2(t *testing.T) {
	type args struct {
		M [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[][]int{
					{1, 1, 0},
					{1, 1, 0},
					{0, 0, 1},
				},
			},
			2,
		},
		{
			"2",
			args{
				[][]int{
					{1, 1, 0},
					{1, 1, 1},
					{0, 1, 1},
				},
			},
			1,
		},
		{
			"3",
			args{
				[][]int{
					{1, 0, 0, 1},
					{0, 1, 1, 0},
					{0, 1, 1, 1},
					{1, 0, 1, 1},
				},
			},
			1,
		},
		{
			"4",
			args{
				[][]int{
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
				},
			},
			1,
		},
		{
			"5",
			args{
				[][]int{
					{1, 1, 1, 0},
					{1, 1, 0, 0},
					{1, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum2(tt.args.M); got != tt.want {
				t.Errorf("findCircleNum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

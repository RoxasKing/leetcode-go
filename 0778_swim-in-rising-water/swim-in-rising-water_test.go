package main

import (
	"testing"
)

func Test_swimInWater(t *testing.T) {
	type args struct {
		grid [][]int
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
					{0, 2},
					{1, 3},
				},
			},
			3,
		},
		{
			"2",
			args{
				[][]int{
					{0, 1, 2, 3, 4},
					{24, 23, 22, 21, 5},
					{12, 13, 14, 15, 16},
					{11, 17, 18, 19, 20},
					{10, 9, 8, 7, 6},
				},
			},
			16,
		},
		{
			"3",
			args{
				[][]int{
					{10, 12, 4, 6},
					{9, 11, 3, 5},
					{1, 7, 13, 8},
					{2, 0, 15, 14},
				},
			},
			14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swimInWater(tt.args.grid); got != tt.want {
				t.Errorf("swimInWater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swimInWater2(t *testing.T) {
	type args struct {
		grid [][]int
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
					{0, 2},
					{1, 3},
				},
			},
			3,
		},
		{
			"2",
			args{
				[][]int{
					{0, 1, 2, 3, 4},
					{24, 23, 22, 21, 5},
					{12, 13, 14, 15, 16},
					{11, 17, 18, 19, 20},
					{10, 9, 8, 7, 6},
				},
			},
			16,
		},
		{
			"3",
			args{
				[][]int{
					{10, 12, 4, 6},
					{9, 11, 3, 5},
					{1, 7, 13, 8},
					{2, 0, 15, 14},
				},
			},
			14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swimInWater2(tt.args.grid); got != tt.want {
				t.Errorf("swimInWater2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swimInWater3(t *testing.T) {
	type args struct {
		grid [][]int
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
					{0, 2},
					{1, 3},
				},
			},
			3,
		},
		{
			"2",
			args{
				[][]int{
					{0, 1, 2, 3, 4},
					{24, 23, 22, 21, 5},
					{12, 13, 14, 15, 16},
					{11, 17, 18, 19, 20},
					{10, 9, 8, 7, 6},
				},
			},
			16,
		},
		{
			"3",
			args{
				[][]int{
					{10, 12, 4, 6},
					{9, 11, 3, 5},
					{1, 7, 13, 8},
					{2, 0, 15, 14},
				},
			},
			14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swimInWater3(tt.args.grid); got != tt.want {
				t.Errorf("swimInWater3() = %v, want %v", got, tt.want)
			}
		})
	}
}

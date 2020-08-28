package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1",
			args{
				[][]int{
					{1, 1},
				},
				2,
			},
			false,
		},
		{
			"2",
			args{
				[][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				3,
			},
			true,
		},
		{
			"3",
			args{
				[][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				13,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

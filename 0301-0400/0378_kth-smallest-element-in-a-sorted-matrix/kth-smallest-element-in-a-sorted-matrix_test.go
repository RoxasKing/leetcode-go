package main

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
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
					{1, 5, 9},
					{10, 11, 13},
					{12, 13, 15},
				},
				8,
			},
			13,
		},
		{
			"2",
			args{
				[][]int{{-5}},
				1,
			},
			-5,
		},
		{
			"3",
			args{
				[][]int{
					{-5, -4},
					{-5, -4},
				},
				2,
			},
			-5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

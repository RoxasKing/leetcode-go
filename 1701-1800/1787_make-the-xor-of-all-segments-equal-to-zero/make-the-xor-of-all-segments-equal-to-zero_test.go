package main

import "testing"

func Test_minChanges(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{1, 2, 0, 3, 0},
				1,
			},
			3,
		},
		{
			"2",
			args{
				[]int{3, 4, 5, 2, 1, 7, 3, 4, 7},
				3,
			},
			3,
		},
		{
			"3",
			args{
				[]int{1, 2, 4, 1, 2, 5, 1, 2, 6},
				3,
			},
			3,
		},
		{
			"4",
			args{
				[]int{11, 20, 3, 18, 26, 30, 20, 7, 3, 0, 25, 9, 19, 21, 3, 23},
				5,
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minChanges(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

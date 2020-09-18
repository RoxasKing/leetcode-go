package main

import "testing"

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				[]int{1, 5, 1, 1, 6, 4},
			},
		},
		{
			"2",
			args{
				[]int{1, 3, 2, 2, 3, 1},
			},
		},
		{
			"3",
			args{
				[]int{1, 1, 2, 1, 2, 2, 1},
			},
		},
		{
			"4",
			args{
				[]int{4, 5, 5, 6},
			},
		},
		{
			"5",
			args{
				[]int{1, 2, 2, 1, 2, 1, 1, 1, 2, 2, 1, 2, 1, 2, 1, 1, 2},
			},
		},
		{
			"6",
			args{
				[]int{1, 1, 2},
			},
		},
		{
			"7",
			args{
				[]int{2, 1},
			},
		},
		{
			"8",
			args{
				[]int{2, 3, 3, 2, 2, 2, 1, 1},
			},
		},
		{
			"9",
			args{
				[]int{4, 5, 5, 5, 5, 6, 6, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
			nums := tt.args.nums
			for i := 1; i < len(nums)-1; i++ {
				if nums[i] > nums[i-1] && nums[i] > nums[i+1] ||
					nums[i] < nums[i-1] && nums[i] < nums[i+1] {
					continue
				}
				t.Error("coinChange() test fail!")
				break
			}
		})
	}
}

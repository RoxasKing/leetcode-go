package main

import (
	"reflect"
	"testing"
)

func Test_maxNumber(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				[]int{3, 4, 6, 5},
				[]int{9, 1, 2, 5, 8, 3},
				5,
			},
			[]int{9, 8, 6, 5, 3},
		},
		{
			"2",
			args{
				[]int{6, 7},
				[]int{6, 0, 4},
				5,
			},
			[]int{6, 7, 6, 0, 4},
		},
		{
			"3",
			args{
				[]int{3, 9},
				[]int{8, 9},
				3,
			},
			[]int{9, 8, 9},
		},
		{
			"4",
			args{
				[]int{8, 9},
				[]int{3, 9},
				3,
			},
			[]int{9, 8, 9},
		},
		{
			"5",
			args{
				[]int{6, 6, 8},
				[]int{5, 0, 9},
				3,
			},
			[]int{9, 6, 8},
		},
		{
			"6",
			args{
				[]int{2, 5, 6, 4, 4, 0},
				[]int{7, 3, 8, 0, 6, 5, 7, 6, 2},
				15,
			},
			[]int{7, 3, 8, 2, 5, 6, 4, 4, 0, 6, 5, 7, 6, 2, 0},
		},
		{
			"7",
			args{
				[]int{2, 1, 7, 8, 0, 1, 7, 3, 5, 8, 9, 0, 0, 7, 0, 2, 2, 7, 3, 5, 5},
				[]int{2, 6, 2, 0, 1, 0, 5, 4, 5, 5, 3, 3, 3, 4},
				35,
			},
			[]int{2, 6, 2, 2, 1, 7, 8, 0, 1, 7, 3, 5, 8, 9, 0, 1, 0, 5, 4, 5, 5, 3, 3, 3, 4, 0, 0, 7, 0, 2, 2, 7, 3, 5, 5},
		},
		{
			"8",
			args{
				[]int{},
				[]int{1},
				1,
			},
			[]int{1},
		},
		{
			"9",
			args{
				[]int{},
				[]int{},
				0,
			},
			[]int{},
		},
		{
			"10",
			args{
				[]int{3, 8, 5, 3, 4},
				[]int{8, 7, 3, 6, 8},
				5,
			},
			[]int{8, 8, 8, 5, 4},
		},
		{
			"11",
			args{
				[]int{4, 3, 3, 0, 6},
				[]int{2, 7, 3, 8, 8, 4, 1, 3, 2, 5, 7, 1, 4, 6, 3, 6, 4, 4, 0},
				24,
			},
			[]int{4, 3, 3, 2, 7, 3, 8, 8, 4, 1, 3, 2, 5, 7, 1, 4, 6, 3, 6, 4, 4, 0, 6, 0},
		},
		{
			"12",
			args{
				[]int{9, 5, 6, 2, 4, 3, 6, 2},
				[]int{5, 7, 6, 2, 2, 1, 3, 0, 2, 8, 9, 7, 7, 3, 2, 2, 9, 4, 5, 1},
				28,
			},
			[]int{9, 5, 7, 6, 5, 6, 2, 4, 3, 6, 2, 2, 2, 1, 3, 0, 2, 8, 9, 7, 7, 3, 2, 2, 9, 4, 5, 1},
		},
		{
			"13",
			args{
				[]int{6, 7, 5},
				[]int{4, 8, 1},
				3,
			},
			[]int{8, 7, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumber(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

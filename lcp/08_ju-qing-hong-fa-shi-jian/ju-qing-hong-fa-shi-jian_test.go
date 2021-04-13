package main

import (
	"reflect"
	"testing"
)

func Test_getTriggerTime(t *testing.T) {
	type args struct {
		increase     [][]int
		requirements [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				[][]int{{2, 8, 4}, {2, 5, 0}, {10, 9, 8}},
				[][]int{{2, 11, 3}, {15, 10, 7}, {9, 17, 12}, {8, 1, 14}},
			},
			[]int{2, -1, 3, -1},
		},
		{
			"2",
			args{
				[][]int{{0, 4, 5}, {4, 8, 8}, {8, 6, 1}, {10, 10, 0}},
				[][]int{{12, 11, 16}, {20, 2, 6}, {9, 2, 6}, {10, 18, 3}, {8, 14, 9}},
			},
			[]int{-1, 4, 3, 3, 3},
		},
		{
			"3",
			args{
				[][]int{{1, 1, 1}},
				[][]int{{0, 0, 0}},
			},
			[]int{0},
		},
		{
			"4",
			args{
				[][]int{{1, 1, 1}},
				[][]int{{0, 0, 1}},
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTriggerTime(tt.args.increase, tt.args.requirements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTriggerTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTriggerTime2(t *testing.T) {
	type args struct {
		increase     [][]int
		requirements [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				[][]int{{2, 8, 4}, {2, 5, 0}, {10, 9, 8}},
				[][]int{{2, 11, 3}, {15, 10, 7}, {9, 17, 12}, {8, 1, 14}},
			},
			[]int{2, -1, 3, -1},
		},
		{
			"2",
			args{
				[][]int{{0, 4, 5}, {4, 8, 8}, {8, 6, 1}, {10, 10, 0}},
				[][]int{{12, 11, 16}, {20, 2, 6}, {9, 2, 6}, {10, 18, 3}, {8, 14, 9}},
			},
			[]int{-1, 4, 3, 3, 3},
		},
		{
			"3",
			args{
				[][]int{{1, 1, 1}},
				[][]int{{0, 0, 0}},
			},
			[]int{0},
		},
		{
			"4",
			args{
				[][]int{{1, 1, 1}},
				[][]int{{0, 0, 1}},
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTriggerTime2(tt.args.increase, tt.args.requirements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTriggerTime2() = %v, want %v", got, tt.want)
			}
		})
	}
}

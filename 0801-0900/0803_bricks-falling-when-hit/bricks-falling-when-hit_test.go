package main

import (
	"reflect"
	"testing"
)

func Test_hitBricks(t *testing.T) {
	type args struct {
		grid [][]int
		hits [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}, [][]int{{1, 0}}}, []int{2}},
		{"2", args{[][]int{{1, 0, 0, 0}, {1, 1, 0, 0}}, [][]int{{1, 1}, {1, 0}}}, []int{0, 0}},
		{"3", args{[][]int{{1, 0, 0, 0}, {1, 1, 1, 1}}, [][]int{{0, 0}}}, []int{4}},
		{
			"4",
			args{
				[][]int{
					{1, 0, 0, 0},
					{1, 1, 1, 0},
					{0, 0, 1, 0},
					{0, 0, 1, 0},
				},
				[][]int{{1, 0}, {1, 2}},
			},
			[]int{4, 0},
		},
		{"5", args{[][]int{{0, 0, 1, 0, 0}, {1, 1, 1, 1, 1}}, [][]int{{0, 2}}}, []int{5}},
		{
			"6",
			args{
				[][]int{
					{1, 0, 0, 0},
					{1, 1, 1, 0},
					{0, 0, 1, 1},
					{0, 0, 1, 1},
				},
				[][]int{{1, 0}, {1, 2}},
			},
			[]int{6, 0},
		},
		{
			"7",
			args{
				[][]int{
					{1, 0, 1, 0},
					{1, 1, 1, 1},
					{0, 0, 0, 1},
					{1, 1, 1, 1},
				},
				[][]int{{1, 0}, {1, 2}},
			},
			[]int{0, 7},
		},
		{
			"8",
			args{[][]int{
				{0, 1, 1, 1, 1},
				{1, 1, 1, 1, 0},
				{1, 1, 1, 1, 0},
				{0, 0, 1, 1, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
				[][]int{{6, 0}, {1, 0}, {4, 3}, {1, 2}, {7, 1}, {6, 3}, {5, 2}, {5, 1}, {2, 4}, {4, 4}, {7, 3}},
			},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hitBricks(tt.args.grid, tt.args.hits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hitBricks() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{[]int{1, 2, 3}},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permute2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{[]int{1, 2, 3}},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_findMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 2},
					},
				},
			},
			[]int{2},
		},
		{
			"2",
			args{
				&TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 2},
					},
				},
			},
			[]int{1, 2},
		},
		{
			"3",
			args{
				&TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			[]int{1, 2},
		},
		{
			"4",
			args{
				nil,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

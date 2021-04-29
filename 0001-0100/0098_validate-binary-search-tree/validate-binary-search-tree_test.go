package main

import (
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				&TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3}},
			},
			true,
		},
		{
			"1",
			args{
				&TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 6},
					},
				},
			},
			false,
		},
		{
			"2",
			args{
				&TreeNode{
					Val:  10,
					Left: &TreeNode{Val: 5},
					Right: &TreeNode{
						Val:   15,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 20},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import (
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{
				&TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			true,
		},
		{
			"",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   3,
							Left:  &TreeNode{Val: 4},
							Right: &TreeNode{Val: 4},
						},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 2},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBalanced2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{
				&TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			true,
		},
		{
			"",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   3,
							Left:  &TreeNode{Val: 4},
							Right: &TreeNode{Val: 4},
						},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 2},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced2(tt.args.root); got != tt.want {
				t.Errorf("isBalanced2() = %v, want %v", got, tt.want)
			}
		})
	}
}

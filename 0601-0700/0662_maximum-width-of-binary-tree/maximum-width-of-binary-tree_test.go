package main

import "testing"

func Test_widthOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
			},
			4,
		},
		{
			"2",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 3},
					},
				},
			},
			2,
		},
		{
			"3",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 5},
					},
					Right: &TreeNode{Val: 2},
				},
			},
			2,
		},
		{
			"4",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  5,
							Left: &TreeNode{Val: 6},
						},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val:   9,
							Right: &TreeNode{Val: 7},
						},
					},
				},
			},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

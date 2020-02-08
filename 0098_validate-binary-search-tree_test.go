package leetcode

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
			"",
			args{
				&TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
			},
			true,
		},
		{
			"",
			args{
				&TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
						Right: &TreeNode{Val: 6, Left: nil, Right: nil},
					},
				},
			},
			false,
		},
		{
			"",
			args{
				&TreeNode{
					Val:  10,
					Left: &TreeNode{Val: 5, Left: nil, Right: nil},
					Right: &TreeNode{
						Val:   15,
						Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
						Right: &TreeNode{Val: 20, Left: nil, Right: nil},
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

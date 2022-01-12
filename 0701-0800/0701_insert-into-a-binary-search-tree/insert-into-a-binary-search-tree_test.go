package main

import "testing"

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 7},
				},
				5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertIntoBST(tt.args.root, tt.args.val)
			node := got
			preVal := -1 << 31
			for node != nil {
				if node.Left != nil {
					pre := node.Left
					for pre.Right != nil && pre.Right != node {
						pre = pre.Right
					}
					if pre.Right != node {
						pre.Right = node
						node = node.Left
						continue
					}
					pre.Right = nil
				}
				if node.Val < preVal {
					t.Error("insertIntoBST() test fail", got)
				}
				preVal = node.Val
				node = node.Right
			}
		})
	}
}

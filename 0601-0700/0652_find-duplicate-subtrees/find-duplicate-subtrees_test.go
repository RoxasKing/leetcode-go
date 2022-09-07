package main

import (
	"fmt"
	"testing"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}
	t5 := &TreeNode{Val: 2}
	t6 := &TreeNode{Val: 4}
	t7 := &TreeNode{Val: 4}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t3.Left = t5
	t3.Right = t6
	t5.Left = t7

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{"1", args{t1}, []*TreeNode{t2, t4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicateSubtrees(tt.args.root)
			ha, hb := map[string]struct{}{}, map[string]struct{}{}
			for _, t := range got {
				ha[serialization(t)] = struct{}{}
			}
			for _, t := range tt.want {
				hb[serialization(t)] = struct{}{}
			}
			for s := range ha {
				if _, ok := hb[s]; !ok {
					t.Errorf("findDuplicateSubtrees() test failed\n")
				}
			}
			for s := range hb {
				if _, ok := ha[s]; !ok {
					t.Errorf("findDuplicateSubtrees() test failed\n")
				}
			}
		})
	}
}

func serialization(node *TreeNode) string {
	if node == nil {
		return ""
	}
	return fmt.Sprintf("%d(%s)(%s)", node.Val, serialization(node.Left), serialization(node.Right))
}

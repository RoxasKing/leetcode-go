package codinginterviews

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	node1 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 8}
	node4 := &TreeNode{Val: 0}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 7}
	node7 := &TreeNode{Val: 9}
	node8 := &TreeNode{Val: 3}
	node9 := &TreeNode{Val: 5}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Left = node8
	node5.Right = node9
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{node1, node2, node3}, node1},
		{"2", args{node1, node2, node5}, node2},
		{"3", args{node1, node4, node9}, node2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestor2(t *testing.T) {
	node1 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 8}
	node4 := &TreeNode{Val: 0}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 7}
	node7 := &TreeNode{Val: 9}
	node8 := &TreeNode{Val: 3}
	node9 := &TreeNode{Val: 5}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Left = node8
	node5.Right = node9
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{node1, node2, node3}, node1},
		{"2", args{node1, node2, node5}, node2},
		{"3", args{node1, node4, node9}, node2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor2(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor2() = %v, want %v", got, tt.want)
			}
		})
	}
}

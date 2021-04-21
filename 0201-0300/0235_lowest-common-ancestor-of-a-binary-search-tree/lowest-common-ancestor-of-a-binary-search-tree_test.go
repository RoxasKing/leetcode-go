package main

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	n0 := &TreeNode{Val: 6}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 8}
	n0.Left = n1
	n0.Right = n2
	n3 := &TreeNode{Val: 0}
	n4 := &TreeNode{Val: 4}
	n1.Left = n3
	n1.Right = n4
	n5 := &TreeNode{Val: 7}
	n6 := &TreeNode{Val: 9}
	n2.Left = n5
	n2.Right = n6
	n7 := &TreeNode{Val: 3}
	n8 := &TreeNode{Val: 5}
	n4.Left = n7
	n4.Right = n8
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
		{
			"1",
			args{n0, n1, n2},
			n0,
		},
		{
			"2",
			args{n0, n1, n4},
			n1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}

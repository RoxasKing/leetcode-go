package main

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	root := &TreeNode{Val: 3}
	n1 := &TreeNode{Val: 5}
	n2 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 6}
	n4 := &TreeNode{Val: 2}
	n5 := &TreeNode{Val: 0}
	n6 := &TreeNode{Val: 8}
	n7 := &TreeNode{Val: 7}
	n8 := &TreeNode{Val: 4}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	n2.Left = n5
	n2.Right = n6
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
			args{root, n1, n2},
			root,
		},
		{
			"2",
			args{root, n1, n8},
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

func Test_lowestCommonAncestor2(t *testing.T) {
	root := &TreeNode{Val: 3}
	n1 := &TreeNode{Val: 5}
	n2 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 6}
	n4 := &TreeNode{Val: 2}
	n5 := &TreeNode{Val: 0}
	n6 := &TreeNode{Val: 8}
	n7 := &TreeNode{Val: 7}
	n8 := &TreeNode{Val: 4}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	n2.Left = n5
	n2.Right = n6
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
			args{root, n1, n2},
			root,
		},
		{
			"2",
			args{root, n1, n8},
			n1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor2(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor2() = %v, want %v", got, tt.want)
			}
		})
	}
}

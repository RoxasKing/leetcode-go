package main

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	a1 := &TreeNode{Val: 5}
	a2 := &TreeNode{Val: 3}
	a3 := &TreeNode{Val: 6}
	a4 := &TreeNode{Val: 2}
	a5 := &TreeNode{Val: 4}
	a6 := &TreeNode{Val: 7}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a2.Right = a5
	a3.Right = a6

	b1 := &TreeNode{Val: 5}
	b2 := &TreeNode{Val: 2}
	b3 := &TreeNode{Val: 6}
	b4 := &TreeNode{Val: 4}
	b5 := &TreeNode{Val: 7}
	b1.Left = b2
	b1.Right = b3
	b2.Right = b4
	b3.Right = b5
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{a1, 3}, b1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

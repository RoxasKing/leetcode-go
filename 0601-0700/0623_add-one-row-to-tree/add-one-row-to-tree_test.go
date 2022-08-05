package main

import (
	"reflect"
	"testing"
)

func Test_addOneRow(t *testing.T) {
	a1 := &TreeNode{Val: 4}
	a2 := &TreeNode{Val: 2}
	a3 := &TreeNode{Val: 3}
	a4 := &TreeNode{Val: 1}
	a1.Left = a2
	a2.Left, a2.Right = a3, a4

	b1 := &TreeNode{Val: 4}
	b2 := &TreeNode{Val: 2}
	b3 := &TreeNode{Val: 1}
	b4 := &TreeNode{Val: 1}
	b5 := &TreeNode{Val: 3}
	b6 := &TreeNode{Val: 1}
	b1.Left = b2
	b2.Left, b2.Right = b3, b4
	b3.Left = b5
	b4.Right = b6

	type args struct {
		root  *TreeNode
		val   int
		depth int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{a1, 1, 3}, b1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.args.root, tt.args.val, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_insertIntoMaxTree(t *testing.T) {
	a1 := &TreeNode{Val: 5}
	a2 := &TreeNode{Val: 2}
	a3 := &TreeNode{Val: 3}
	a4 := &TreeNode{Val: 1}
	a1.Left = a2
	a1.Right = a3
	a2.Right = a4

	b1 := &TreeNode{Val: 5}
	b2 := &TreeNode{Val: 2}
	b3 := &TreeNode{Val: 4}
	b4 := &TreeNode{Val: 1}
	b5 := &TreeNode{Val: 3}
	b1.Left = b2
	b1.Right = b3
	b2.Right = b4
	b3.Left = b5

	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{a1, 4}, b1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoMaxTree(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoMaxTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

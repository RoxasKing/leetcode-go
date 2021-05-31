package main

import (
	"reflect"
	"testing"
)

func Test_subtreeWithAllDeepest(t *testing.T) {
	a1 := &TreeNode{Val: 3}
	a2 := &TreeNode{Val: 5}
	a3 := &TreeNode{Val: 1}
	a4 := &TreeNode{Val: 6}
	a5 := &TreeNode{Val: 2}
	a6 := &TreeNode{Val: 0}
	a7 := &TreeNode{Val: 8}
	a8 := &TreeNode{Val: 7}
	a9 := &TreeNode{Val: 4}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a2.Right = a5
	a3.Left = a6
	a3.Right = a7
	a5.Left = a8
	a5.Right = a9

	b1 := &TreeNode{Val: 1}

	c1 := &TreeNode{Val: 0}
	c2 := &TreeNode{Val: 1}
	c3 := &TreeNode{Val: 3}
	c4 := &TreeNode{Val: 2}
	c1.Left = c2
	c1.Right = c3
	c2.Right = c4

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{a1}, a5},
		{"2", args{b1}, b1},
		{"3", args{c1}, c4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtreeWithAllDeepest(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtreeWithAllDeepest() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_verticalTraversal(t *testing.T) {
	a1 := &TreeNode{Val: 3}
	a2 := &TreeNode{Val: 9}
	a3 := &TreeNode{Val: 20}
	a4 := &TreeNode{Val: 15}
	a5 := &TreeNode{Val: 7}
	a1.Left = a2
	a1.Right = a3
	a3.Left = a4
	a3.Right = a5

	b1 := &TreeNode{Val: 1}
	b2 := &TreeNode{Val: 2}
	b3 := &TreeNode{Val: 3}
	b4 := &TreeNode{Val: 4}
	b5 := &TreeNode{Val: 5}
	b6 := &TreeNode{Val: 6}
	b7 := &TreeNode{Val: 7}
	b1.Left = b2
	b1.Right = b3
	b2.Left = b4
	b2.Right = b5
	b3.Left = b6
	b3.Right = b7

	c1 := &TreeNode{Val: 1}
	c2 := &TreeNode{Val: 2}
	c3 := &TreeNode{Val: 3}
	c4 := &TreeNode{Val: 4}
	c5 := &TreeNode{Val: 6}
	c6 := &TreeNode{Val: 5}
	c7 := &TreeNode{Val: 7}
	c1.Left = c2
	c1.Right = c3
	c2.Left = c4
	c2.Right = c5
	c3.Left = c6
	c3.Right = c7

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{a1}, [][]int{{9}, {3, 15}, {20}, {7}}},
		{"2", args{b1}, [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}},
		{"3", args{c1}, [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

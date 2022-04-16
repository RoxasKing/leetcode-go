package main

import (
	"reflect"
	"testing"
)

func Test_trimBST(t *testing.T) {
	a1 := &TreeNode{Val: 1}
	a2 := &TreeNode{Val: 0}
	a3 := &TreeNode{Val: 2}
	a1.Left = a2
	a1.Right = a3

	a0 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	b1 := &TreeNode{Val: 3}
	b2 := &TreeNode{Val: 0}
	b3 := &TreeNode{Val: 4}
	b4 := &TreeNode{Val: 2}
	b5 := &TreeNode{Val: 1}
	b1.Left = b2
	b1.Right = b3
	b2.Right = b4
	b4.Left = b5

	b0 := &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}

	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{a1, 1, 2}, a0},
		{"2", args{b1, 1, 3}, b0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBST(tt.args.root, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

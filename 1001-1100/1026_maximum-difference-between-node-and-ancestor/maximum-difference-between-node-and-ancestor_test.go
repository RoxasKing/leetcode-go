package main

import "testing"

func Test_maxAncestorDiff(t *testing.T) {
	a1 := &TreeNode{Val: 8}
	a2 := &TreeNode{Val: 3}
	a3 := &TreeNode{Val: 10}
	a4 := &TreeNode{Val: 1}
	a5 := &TreeNode{Val: 6}
	a6 := &TreeNode{Val: 14}
	a7 := &TreeNode{Val: 4}
	a8 := &TreeNode{Val: 7}
	a9 := &TreeNode{Val: 13}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a2.Right = a5
	a3.Right = a6
	a5.Left = a7
	a5.Right = a8
	a6.Left = a9

	b1 := &TreeNode{Val: 1}
	b2 := &TreeNode{Val: 2}
	b3 := &TreeNode{Val: 0}
	b4 := &TreeNode{Val: 4}
	b1.Right = b2
	b2.Right = b3
	b3.Left = b4

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1}, 7},
		{"2", args{b1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAncestorDiff(tt.args.root); got != tt.want {
				t.Errorf("maxAncestorDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_sumRootToLeaf(t *testing.T) {
	a1 := &TreeNode{Val: 1}
	a2 := &TreeNode{Val: 0}
	a3 := &TreeNode{Val: 1}
	a4 := &TreeNode{Val: 0}
	a5 := &TreeNode{Val: 1}
	a6 := &TreeNode{Val: 0}
	a7 := &TreeNode{Val: 1}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a2.Right = a5
	a3.Left = a6
	a3.Right = a7

	b1 := &TreeNode{Val: 0}

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1}, 22},
		{"2", args{b1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumRootToLeaf(tt.args.root); got != tt.want {
				t.Errorf("sumRootToLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

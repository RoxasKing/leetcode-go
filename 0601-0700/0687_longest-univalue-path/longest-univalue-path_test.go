package main

import "testing"

func Test_longestUnivaluePath(t *testing.T) {
	a1 := &TreeNode{Val: 5}
	a2 := &TreeNode{Val: 4}
	a3 := &TreeNode{Val: 5}
	a4 := &TreeNode{Val: 1}
	a5 := &TreeNode{Val: 1}
	a6 := &TreeNode{Val: 5}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a2.Right = a5
	a3.Right = a6

	b1 := &TreeNode{Val: 1}
	b2 := &TreeNode{Val: 4}
	b3 := &TreeNode{Val: 5}
	b4 := &TreeNode{Val: 4}
	b5 := &TreeNode{Val: 4}
	b6 := &TreeNode{Val: 5}
	b1.Left = b2
	b1.Right = b3
	b2.Left = b4
	b2.Right = b5
	b3.Right = b6

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1}, 2},
		{"2", args{b1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestUnivaluePath(tt.args.root); got != tt.want {
				t.Errorf("longestUnivaluePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

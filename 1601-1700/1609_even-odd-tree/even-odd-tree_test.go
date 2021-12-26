package main

import "testing"

func Test_isEvenOddTree(t *testing.T) {
	a0 := &TreeNode{Val: 1}
	a1 := &TreeNode{Val: 10}
	a2 := &TreeNode{Val: 4}
	a3 := &TreeNode{Val: 3}
	a4 := &TreeNode{Val: 7}
	a5 := &TreeNode{Val: 9}
	a6 := &TreeNode{Val: 12}
	a7 := &TreeNode{Val: 8}
	a8 := &TreeNode{Val: 6}
	a9 := &TreeNode{Val: 2}
	a0.Left = a1
	a0.Right = a2
	a1.Left = a3
	a2.Left = a4
	a2.Right = a5
	a3.Left = a6
	a3.Right = a7
	a4.Left = a8
	a5.Right = a9

	b0 := &TreeNode{Val: 5}
	b1 := &TreeNode{Val: 4}
	b2 := &TreeNode{Val: 2}
	b3 := &TreeNode{Val: 3}
	b4 := &TreeNode{Val: 3}
	b5 := &TreeNode{Val: 7}
	b0.Left = b1
	b0.Right = b2
	b1.Left = b3
	b1.Right = b4
	b2.Left = b5

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{a0}, true},
		{"2", args{b0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEvenOddTree(tt.args.root); got != tt.want {
				t.Errorf("isEvenOddTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

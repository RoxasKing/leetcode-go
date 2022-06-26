package main

import "testing"

func Test_findBottomLeftValue(t *testing.T) {
	a1 := &TreeNode{Val: 2}
	a2 := &TreeNode{Val: 1}
	a3 := &TreeNode{Val: 3}
	a1.Left = a2
	a1.Right = a3

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
	b3.Left = b5
	b3.Right = b6
	b5.Left = b7

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1}, 1},
		{"2", args{b1}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_goodNodes(t *testing.T) {
	a1 := &TreeNode{Val: 3}
	a2 := &TreeNode{Val: 1}
	a3 := &TreeNode{Val: 4}
	a4 := &TreeNode{Val: 3}
	a5 := &TreeNode{Val: 1}
	a6 := &TreeNode{Val: 5}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a3.Left = a5
	a3.Right = a6

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodNodes(tt.args.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

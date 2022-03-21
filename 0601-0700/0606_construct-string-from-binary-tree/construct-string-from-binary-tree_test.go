package main

import "testing"

func Test_tree2str(t *testing.T) {
	a1 := &TreeNode{Val: 1}
	a2 := &TreeNode{Val: 2}
	a3 := &TreeNode{Val: 3}
	a4 := &TreeNode{Val: 4}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4

	b1 := &TreeNode{Val: 1}
	b2 := &TreeNode{Val: 2}
	b3 := &TreeNode{Val: 3}
	b4 := &TreeNode{Val: 4}
	b1.Left = b2
	b1.Right = b3
	b2.Right = b4

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{a1}, "1(2(4))(3)"},
		{"2", args{b1}, "1(2()(4))(3)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.root); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}

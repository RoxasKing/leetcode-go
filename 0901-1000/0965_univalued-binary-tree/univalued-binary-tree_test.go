package main

import "testing"

func Test_isUnivalTree(t *testing.T) {
	t1 := &TreeNode{Val: 2}
	t2 := &TreeNode{Val: 1}
	t3 := &TreeNode{Val: 1}
	t4 := &TreeNode{Val: 1}
	t5 := &TreeNode{Val: 1}
	t1.Left = t2
	t2.Left = t3
	t2.Left = t4
	t2.Right = t5

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{t1}, false},
		{"2", args{t2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnivalTree(tt.args.root); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

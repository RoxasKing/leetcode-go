package main

import "testing"

func Test_pathSum(t *testing.T) {
	a1 := &TreeNode{Val: 10}
	a2 := &TreeNode{Val: 5}
	a3 := &TreeNode{Val: -3}
	a4 := &TreeNode{Val: 3}
	a5 := &TreeNode{Val: 2}
	a6 := &TreeNode{Val: 11}
	a7 := &TreeNode{Val: 3}
	a8 := &TreeNode{Val: -2}
	a9 := &TreeNode{Val: 1}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a2.Right = a5
	a3.Right = a6
	a4.Left = a7
	a4.Right = a8
	a5.Right = a9
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1, 8}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

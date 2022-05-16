package main

import "testing"

func Test_deepestLeavesSum(t *testing.T) {
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}
	t5 := &TreeNode{Val: 5}
	t6 := &TreeNode{Val: 6}
	t7 := &TreeNode{Val: 7}
	t8 := &TreeNode{Val: 8}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t3.Right = t6
	t4.Left = t7
	t6.Right = t8

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{t1}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepestLeavesSum(tt.args.root); got != tt.want {
				t.Errorf("deepestLeavesSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

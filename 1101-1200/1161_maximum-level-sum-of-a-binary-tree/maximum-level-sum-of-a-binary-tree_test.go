package main

import "testing"

func Test_maxLevelSum(t *testing.T) {
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 7}
	t3 := &TreeNode{Val: 0}
	t4 := &TreeNode{Val: 7}
	t5 := &TreeNode{Val: -8}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{t1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLevelSum(tt.args.root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_isCousins(t *testing.T) {
	type args struct {
		root *TreeNode
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			}, 4, 3},
			false,
		},
		{
			"2",
			args{&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 5},
				},
			}, 5, 4},
			true,
		},
		{
			"3",
			args{&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			}, 2, 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins(tt.args.root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}

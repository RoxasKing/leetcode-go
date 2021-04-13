package main

import (
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				&TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
				&TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			true,
		},
		{
			"2",
			args{
				&TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
				&TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
			},
			false,
		},
		{
			"3",
			args{
				&TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 1},
				},
				&TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

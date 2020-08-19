package main

import (
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val:  6,
						Left: &TreeNode{Val: 7},
						Right: &TreeNode{
							Val:   8,
							Right: &TreeNode{Val: 9},
						},
					},
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

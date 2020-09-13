package main

import (
	"testing"
)

func Test_maxPathSum(t *testing.T) {
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
					Val:  -10,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

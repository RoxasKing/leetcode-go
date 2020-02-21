package leetcode

import (
	"testing"
)

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				&TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepth2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				&TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth2(tt.args.root); got != tt.want {
				t.Errorf("minDepth2() = %v, want %v", got, tt.want)
			}
		})
	}
}

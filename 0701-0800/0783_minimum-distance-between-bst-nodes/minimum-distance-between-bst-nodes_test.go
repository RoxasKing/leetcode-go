package main

import "testing"

func Test_minDiffInBST(t *testing.T) {
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
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 6},
				},
			},
			1,
		},
		{
			"2",
			args{
				&TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   48,
						Left:  &TreeNode{Val: 12},
						Right: &TreeNode{Val: 49},
					},
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDiffInBST(tt.args.root); got != tt.want {
				t.Errorf("minDiffInBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

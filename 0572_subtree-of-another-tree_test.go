package main

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		s *TreeNode
		t *TreeNode
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
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 2},
					},
					Right: &TreeNode{Val: 5},
				},
				&TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
			},
			true,
		},
		{
			"2",
			args{
				&TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 1},
						Right: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 0},
						},
					},
					Right: &TreeNode{Val: 5},
				},
				&TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}

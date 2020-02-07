package My_LeetCode_In_Go

import (
	. "My_LeetCode_In_Go/util/tree"
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
			"",
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
			"",
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
			"",
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

func Test_isSameTree_Stack(t *testing.T) {
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
			"",
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
			"",
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
			"",
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
			if got := isSameTree_Stack(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree_Stack() = %v, want %v", got, tt.want)
			}
		})
	}
}

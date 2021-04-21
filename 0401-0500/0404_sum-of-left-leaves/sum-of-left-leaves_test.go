package main

import (
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
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
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			24,
		},
		{
			"2",
			args{
				&TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   1,
							Left:  &TreeNode{Val: 5},
							Right: &TreeNode{Val: 1},
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   3,
							Right: &TreeNode{Val: 6},
						},
						Right: &TreeNode{
							Val:   -1,
							Right: &TreeNode{Val: 8},
						},
					},
				},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfLeftLeaves2(t *testing.T) {
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
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			24,
		},
		{
			"2",
			args{
				&TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   1,
							Left:  &TreeNode{Val: 5},
							Right: &TreeNode{Val: 1},
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   3,
							Right: &TreeNode{Val: 6},
						},
						Right: &TreeNode{
							Val:   -1,
							Right: &TreeNode{Val: 8},
						},
					},
				},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves2(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves2() = %v, want %v", got, tt.want)
			}
		})
	}
}

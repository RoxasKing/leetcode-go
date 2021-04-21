package main

import (
	"testing"
)

func Test_rob(t *testing.T) {
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
			args{&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 1},
				},
			}},
			7,
		},
		{
			"2",
			args{&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 1},
				},
			}},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.root); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

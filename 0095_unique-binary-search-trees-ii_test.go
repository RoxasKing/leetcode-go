package main

import (
	"reflect"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			"1",
			args{3},
			[]*TreeNode{
				{
					Val: 1,
					Right: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 3},
					},
				},
				{
					Val: 1,
					Right: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 2},
					},
				},
				{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				{
					Val: 3,
					Left: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 2},
					},
				},
				{
					Val: 3,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 1},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

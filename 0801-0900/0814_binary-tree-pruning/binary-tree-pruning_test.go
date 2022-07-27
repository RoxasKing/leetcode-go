package main

import (
	"reflect"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"1",
			args{&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 1},
				},
			}},
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   0,
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			"2",
			args{&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 0},
				},
				Right: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 1},
				},
			}},
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   0,
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			"3",
			args{&TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 0},
				},
			}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

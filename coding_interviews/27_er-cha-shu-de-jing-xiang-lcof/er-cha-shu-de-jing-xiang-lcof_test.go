package main

import (
	"reflect"
	"testing"
)

func Test_mirrorTree(t *testing.T) {
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
			args{
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 9},
					},
				},
			},
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mirrorTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_buildTree0106(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"",
			args{
				[]int{9, 3, 15, 20, 7},
				[]int{9, 15, 7, 20, 3},
			},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree2(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree0106() = %v, want %v", got, tt.want)
			}
		})
	}
}

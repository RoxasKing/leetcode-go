package My_LeetCode_In_Go

import (
	. "My_LeetCode_In_Go/util/tree"
	"reflect"
	"testing"
)

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"",
			args{
				&TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}},
				},
			},
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			},
		},
		{
			"",
			args{
				&TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}},
				},
			},
			&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("recoverTree() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}

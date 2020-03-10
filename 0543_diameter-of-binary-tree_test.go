package leetcode

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{&TreeNode{Val: 1}}, 0},
		{
			"",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{Val: 3},
				},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

package codinginterviews

import "testing"

func Test_isSubStructure(t *testing.T) {
	type args struct {
		A *TreeNode
		B *TreeNode
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
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
				&TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 1},
				},
			},
			false,
		},
		{
			"2",
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
					Val:  4,
					Left: &TreeNode{Val: 1},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

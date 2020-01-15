package My_LeetCode_In_Go

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				&TreeNode{
					2,
					&TreeNode{
						3,
						&TreeNode{4, nil, nil},
						&TreeNode{5, nil, nil},
					},
					&TreeNode{
						6,
						&TreeNode{7, nil, nil},
						&TreeNode{
							8,
							nil,
							&TreeNode{9, nil, nil},
						},
					},
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

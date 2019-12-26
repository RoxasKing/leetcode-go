package My_LeetCode_In_Go

import "testing"

func Test_maxPathSum(t *testing.T) {
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
					-10,
					&TreeNode{9, nil, nil},
					&TreeNode{
						20,
						&TreeNode{15, nil, nil},
						&TreeNode{7, nil, nil},
					},
				},
			},
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

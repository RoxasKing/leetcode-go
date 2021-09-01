package main

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 4},
					},
					Right: &TreeNode{Val: 3},
				},
				2,
			},
			12,
		},
		{
			"2",
			args{
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 9},
					},
					Right: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 2},
					},
				},
				2,
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
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
					Val: 10,
					Left: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 7},
					},
					Right: &TreeNode{
						Val:   15,
						Right: &TreeNode{Val: 18},
					},
				},
				7, 15,
			},
			32,
		},
		{
			"2",
			args{
				&TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:  3,
							Left: &TreeNode{Val: 1},
						},
						Right: &TreeNode{
							Val:  7,
							Left: &TreeNode{Val: 6},
						},
					},
					Right: &TreeNode{
						Val:   15,
						Left:  &TreeNode{Val: 13},
						Right: &TreeNode{Val: 18},
					},
				},
				6, 10,
			},
			23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

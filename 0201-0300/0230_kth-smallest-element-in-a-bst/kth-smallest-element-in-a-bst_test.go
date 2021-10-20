package main

import "testing"

func Test_kthSmallest(t *testing.T) {
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
					Val: 3,
					Left: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 2},
					},
					Right: &TreeNode{Val: 4},
				},
				1,
			},
			1,
		},
		{
			"2",
			args{
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 1},
						},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{Val: 6},
				},
				3,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

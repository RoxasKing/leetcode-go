package codinginterviews

import (
	"testing"
)

func Test_kthLargest(t *testing.T) {
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
			4,
		},
		{
			"1",
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
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kthLargest2(t *testing.T) {
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
			4,
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
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest2(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargest2() = %v, want %v", got, tt.want)
			}
		})
	}
}

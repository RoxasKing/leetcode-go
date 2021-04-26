package main

import (
	"reflect"
	"testing"
)

func Test_increasingBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
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
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val:   8,
							Left:  &TreeNode{Val: 7},
							Right: &TreeNode{Val: 9},
						},
					},
				},
			},
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
									Right: &TreeNode{
										Val: 7,
										Right: &TreeNode{
											Val: 8,
											Right: &TreeNode{
												Val: 9,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			"2",
			args{
				&TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 7},
				},
			},
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 7},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increasingBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

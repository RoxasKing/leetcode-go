package main

import "testing"

func Test_minimalExecTime(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"1",
			args{
				&TreeNode{
					Val:   47,
					Left:  &TreeNode{Val: 74},
					Right: &TreeNode{Val: 31},
				},
			},
			121,
		},
		{
			"2",
			args{
				&TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 21,
						Left: &TreeNode{
							Val:   24,
							Left:  &TreeNode{Val: 27},
							Right: &TreeNode{Val: 26},
						},
					},
				},
			},
			87,
		},
		{
			"3",
			args{
				&TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 3},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 4},
					},
				},
			},
			7.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalExecTime(tt.args.root); got != tt.want {
				t.Errorf("minimalExecTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

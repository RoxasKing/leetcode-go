package main

import "testing"

func Test_minCameraCover(t *testing.T) {
	type args struct {
		root *TreeNode
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
					Left: &TreeNode{
						Left:  &TreeNode{},
						Right: &TreeNode{},
					},
				},
			},
			1,
		},
		{
			"2",
			args{
				&TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{
							Left: &TreeNode{
								Right: &TreeNode{},
							},
						},
					},
				},
			},
			2,
		},
		{
			"3",
			args{
				&TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{
							Left:  &TreeNode{},
							Right: &TreeNode{},
						},
						Right: &TreeNode{
							Right: &TreeNode{},
						},
					},
					Right: &TreeNode{
						Right: &TreeNode{
							Right: &TreeNode{
								Right: &TreeNode{},
							},
						},
					},
				},
			},
			4,
		},
		{"4", args{&TreeNode{}}, 1},
		{
			"5",
			args{
				&TreeNode{
					Left: &TreeNode{
						Right: &TreeNode{
							Left: &TreeNode{
								Right: &TreeNode{
									Left: &TreeNode{},
								},
							},
						},
					},
				},
			},
			2,
		},
		{
			"6",
			args{
				&TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{
							Left: &TreeNode{
								Right: &TreeNode{
									Right: &TreeNode{},
								},
							},
						},
						Right: &TreeNode{
							Right: &TreeNode{},
						},
					},
					Right: &TreeNode{},
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}

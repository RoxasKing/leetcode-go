package main

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			[]int{1, 2, 3},
		},
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 6},
					},
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			[]int{1, 4, 5, 6, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversal2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			[]int{1, 2, 3},
		},
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 6},
					},
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			[]int{1, 4, 5, 6, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversal3(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			[]int{1, 2, 3},
		},
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 6},
					},
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			[]int{1, 4, 5, 6, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal3(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal3() = %v, want %v", got, tt.want)
			}
		})
	}
}

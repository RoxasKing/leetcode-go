package main

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{Val: 3},
				},
			},
			[]string{"1->2->5", "1->3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryTreePaths2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{Val: 3},
				},
			},
			[]string{"1->2->5", "1->3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binaryTreePaths2(tt.args.root)
			memo1 := make(map[string]bool)
			memo2 := make(map[string]bool)
			for _, str := range got {
				memo1[str] = true
			}
			for _, str := range tt.want {
				memo2[str] = true
			}
			for _, str := range got {
				if !memo2[str] {
					t.Errorf("binaryTreePaths2() = %v, want %v", got, tt.want)
				}
			}
			for _, str := range tt.want {
				if !memo1[str] {
					t.Errorf("binaryTreePaths2() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	t1 := &TreeNode{Val: 6}
	t2 := &TreeNode{Val: 3}
	t3 := &TreeNode{Val: 5}
	t4 := &TreeNode{Val: 2}
	t5 := &TreeNode{Val: 0}
	t6 := &TreeNode{Val: 1}
	t1.Left = t2
	t1.Right = t3
	t2.Right = t4
	t3.Left = t5
	t4.Right = t6

	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{[]int{3, 2, 1, 6, 0, 5}}, t1},
		{"2", args{[]int{3, 2, 1}}, t2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

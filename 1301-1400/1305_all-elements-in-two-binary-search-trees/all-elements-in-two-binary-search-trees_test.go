package main

import (
	"reflect"
	"testing"
)

func Test_getAllElements(t *testing.T) {
	a1 := &TreeNode{Val: 2}
	a2 := &TreeNode{Val: 1}
	a3 := &TreeNode{Val: 4}
	a1.Left, a1.Right = a2, a3
	b1 := &TreeNode{Val: 1}
	b2 := &TreeNode{Val: 0}
	b3 := &TreeNode{Val: 3}
	b1.Left, b1.Right = b2, b3

	c1 := &TreeNode{Val: 1}
	c2 := &TreeNode{Val: 8}
	c1.Right = c2
	d1 := &TreeNode{Val: 8}
	d2 := &TreeNode{Val: 1}
	d1.Left = d2

	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{a1, b1}, []int{0, 1, 1, 2, 3, 4}},
		{"2", args{c1, d1}, []int{1, 1, 8, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllElements(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

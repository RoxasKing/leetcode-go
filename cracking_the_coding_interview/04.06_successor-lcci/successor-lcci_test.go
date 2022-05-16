package main

import (
	"reflect"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	a1 := &TreeNode{Val: 2}
	a2 := &TreeNode{Val: 1}
	a3 := &TreeNode{Val: 3}
	a1.Left = a2
	a1.Right = a3

	b1 := &TreeNode{Val: 5}
	b2 := &TreeNode{Val: 3}
	b3 := &TreeNode{Val: 6}
	b4 := &TreeNode{Val: 2}
	b5 := &TreeNode{Val: 4}
	b6 := &TreeNode{Val: 1}
	b1.Left = b2
	b1.Right = b3
	b2.Left = b4
	b2.Right = b5
	b4.Left = b6

	type args struct {
		root *TreeNode
		p    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{a1, a2}, a1},
		{"2", args{b1, b3}, nil},
		{"3", args{b1, b5}, b1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderSuccessor(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

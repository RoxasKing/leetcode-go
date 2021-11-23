package main

import "testing"

func Test_findTilt(t *testing.T) {
	a1 := &TreeNode{Val: 1}
	a2 := &TreeNode{Val: 2}
	a3 := &TreeNode{Val: 3}
	a1.Left = a2
	a1.Right = a3

	b1 := &TreeNode{Val: 4}
	b2 := &TreeNode{Val: 2}
	b3 := &TreeNode{Val: 9}
	b4 := &TreeNode{Val: 3}
	b5 := &TreeNode{Val: 5}
	b6 := &TreeNode{Val: 7}
	b1.Left = b2
	b1.Right = b3
	b2.Left = b4
	b2.Right = b5
	b3.Right = b6

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1}, 1},
		{"2", args{b1}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTilt(tt.args.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}

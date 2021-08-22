package main

import "testing"

func Test_maxProduct(t *testing.T) {
	a1 := &TreeNode{Val: 1}
	a2 := &TreeNode{Val: 2}
	a3 := &TreeNode{Val: 3}
	a4 := &TreeNode{Val: 4}
	a5 := &TreeNode{Val: 5}
	a6 := &TreeNode{Val: 6}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a2.Right = a5
	a3.Right = a6

	b1 := &TreeNode{Val: 1}
	b2 := &TreeNode{Val: 2}
	b3 := &TreeNode{Val: 3}
	b4 := &TreeNode{Val: 4}
	b5 := &TreeNode{Val: 5}
	b6 := &TreeNode{Val: 6}
	b1.Right = b2
	b2.Left = b3
	b2.Right = b4
	b4.Left = b5
	b4.Right = b6

	c1 := &TreeNode{Val: 2}
	c2 := &TreeNode{Val: 3}
	c3 := &TreeNode{Val: 9}
	c4 := &TreeNode{Val: 10}
	c5 := &TreeNode{Val: 7}
	c6 := &TreeNode{Val: 8}
	c7 := &TreeNode{Val: 6}
	c8 := &TreeNode{Val: 5}
	c9 := &TreeNode{Val: 4}
	c10 := &TreeNode{Val: 11}
	c11 := &TreeNode{Val: 1}
	c1.Left = c2
	c1.Right = c3
	c2.Left = c4
	c2.Right = c5
	c3.Left = c6
	c3.Right = c7
	c4.Left = c8
	c4.Right = c9
	c5.Left = c10
	c5.Right = c11

	d1 := &TreeNode{Val: 1}
	d2 := &TreeNode{Val: 1}
	d1.Left = d2

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1}, 110},
		{"2", args{b1}, 90},
		{"3", args{c1}, 1025},
		{"4", args{d1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.root); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_findTarget(t *testing.T) {
	n1 := &TreeNode{Val: 5}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 6}
	n4 := &TreeNode{Val: 2}
	n5 := &TreeNode{Val: 4}
	n6 := &TreeNode{Val: 7}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Right = n6

	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{n1, 9}, true},
		{"2", args{n1, 28}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

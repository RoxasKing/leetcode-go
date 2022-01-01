package main

import "testing"

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				&Node{
					Val: 1,
					Left: &Node{
						Val:   2,
						Left:  &Node{Val: 4},
						Right: &Node{Val: 5},
					},
					Right: &Node{
						Val:   3,
						Left:  &Node{Val: 6},
						Right: &Node{Val: 7},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !check(got) {
				t.Errorf("connect() test failed!")
			}
		})
	}
}

func check(root *Node) bool {
	node := root
	for node != nil {
		if node.Left != nil {
			pre := node.Left
			for pre.Right != nil && pre.Right != node {
				pre = pre.Right
			}
			if pre.Right != node {
				pre.Right = node
				node = node.Left
				continue
			}
			pre.Right = nil
			for l, r := node.Left, node.Right; l != nil && r != nil; l, r = l.Right, r.Left {
				if l.Next != r {
					return false
				}
			}
		}
		node = node.Right
	}
	return true
}

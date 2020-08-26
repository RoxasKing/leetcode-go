package main

import (
	"testing"
)

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
			got := connect(tt.args.root)
			var bfs func(*Node)
			bfs = func(node *Node) {
				if node.Left == nil && node.Right == nil {
					return
				}
				nodeL, nodeR := node.Left, node.Right
				for nodeL != nil && nodeR != nil {
					if nodeL.Next != nodeR {
						t.Errorf("connect() failed")
					}
					nodeL, nodeR = nodeL.Right, nodeR.Left
				}
				bfs(node.Left)
				bfs(node.Right)
			}
			bfs(got)
		})
	}
}

func Test_connect2(t *testing.T) {
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
			got := connect2(tt.args.root)
			var bfs func(*Node)
			bfs = func(node *Node) {
				if node.Left == nil && node.Right == nil {
					return
				}
				nodeL, nodeR := node.Left, node.Right
				for nodeL != nil && nodeR != nil {
					if nodeL.Next != nodeR {
						t.Errorf("connect2() failed")
					}
					nodeL, nodeR = nodeL.Right, nodeR.Left
				}
				bfs(node.Left)
				bfs(node.Right)
			}
			bfs(got)
		})
	}
}

package main

// Difficulty:
// Medium

// Tags:
// Morris Traversal

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
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
				l.Next = r
			}
		}
		node = node.Right
	}
	return root
}

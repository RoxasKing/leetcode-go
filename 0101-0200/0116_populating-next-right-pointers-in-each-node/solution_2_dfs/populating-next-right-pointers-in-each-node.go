package main

// Difficulty:
// Medium

// Tags:
// DFS

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var dfs func(*Node)
	dfs = func(node *Node) {
		for l, r := node.Left, node.Right; l != nil && r != nil; l, r = l.Right, r.Left {
			l.Next = r
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	return root
}

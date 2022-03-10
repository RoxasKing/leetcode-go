package main

// Difficulty:
// Easy

// Tags:
// Stack
// DFS

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	out := []int{}
	if root == nil {
		return out
	}
	for q := []*Node{root}; len(q) > 0; {
		node := q[len(q)-1]
		q = q[:len(q)-1]
		out = append(out, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			q = append(q, node.Children[i])
		}
	}
	return out
}

package main

// Difficulty:
// Eash

// Tags:
// DFS

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	out := []int{}
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, next := range node.Children {
			dfs(next)
		}
		out = append(out, node.Val)
	}
	dfs(root)
	return out
}

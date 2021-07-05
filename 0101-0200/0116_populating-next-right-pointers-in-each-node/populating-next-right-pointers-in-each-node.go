package main

// Tags:
// DFS + Recursion
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	dfs(root)
	return root
}

func dfs(node *Node) {
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

// Morris Traversal
func connect2(root *Node) *Node {
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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

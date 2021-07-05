package main

import "strconv"

// DFS + Recursion
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var out []string
	var dfs func(string, *TreeNode)
	dfs = func(cur string, node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			out = append(out, cur)
			return
		}
		cur += "->"
		if node.Left != nil {
			dfs(cur+strconv.Itoa(node.Left.Val), node.Left)
		}
		if node.Right != nil {
			dfs(cur+strconv.Itoa(node.Right.Val), node.Right)
		}
	}
	dfs(strconv.Itoa(root.Val), root)
	return out
}

// BFS + Queue
func binaryTreePaths2(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	type elem struct {
		node *TreeNode
		path string
	}
	var out []string
	q := []elem{{node: root, path: strconv.Itoa(root.Val)}}
	for len(q) != 0 {
		e := q[0]
		q = q[1:]
		if e.node.Left == nil && e.node.Right == nil {
			out = append(out, e.path)
		}
		if e.node.Left != nil {
			lPath := e.path + "->" + strconv.Itoa(e.node.Left.Val)
			q = append(q, elem{node: e.node.Left, path: lPath})
		}
		if e.node.Right != nil {
			rPath := e.path + "->" + strconv.Itoa(e.node.Right.Val)
			q = append(q, elem{node: e.node.Right, path: rPath})
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

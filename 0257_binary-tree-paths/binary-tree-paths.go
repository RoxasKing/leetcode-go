package main

import "strconv"

/*
  给定一个二叉树，返回所有从根节点到叶子节点的路径。

  说明: 叶子节点是指没有子节点的节点。
*/

// DFS + Recursion
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var out []string
	var dfs func(string, *TreeNode)
	dfs = func(cur string, node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			cur += strconv.Itoa(node.Val)
			out = append(out, cur)
			return
		}
		cur += strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			dfs(cur, node.Left)
		}
		if node.Right != nil {
			dfs(cur, node.Right)
		}
	}
	dfs("", root)
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

package main

/*
  从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

  提示：
    节点总数 <= 1000
*/

// Preorder traversal (VLR) + BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	out := []int{}
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		out = append(out, e.Val)
		if e.Left != nil {
			q = append(q, e.Left)
		}
		if e.Right != nil {
			q = append(q, e.Right)
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

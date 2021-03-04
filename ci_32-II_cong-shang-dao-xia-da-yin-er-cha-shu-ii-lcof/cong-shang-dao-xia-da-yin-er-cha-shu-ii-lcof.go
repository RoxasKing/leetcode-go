package main

/*
  从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

  提示：
    节点总数 <= 1000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	out := [][]int{}
	for len(q) > 0 {
		tmp := make([]int, 0, len(q))
		for _, t := range q {
			tmp = append(tmp, t.Val)
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		q = q[len(tmp):]
		out = append(out, tmp)
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

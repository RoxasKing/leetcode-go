package main

/*
  给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
*/

// BFS + Iteration
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var out [][]int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		cur := make([]int, 0, len(queue))
		for _, node := range queue {
			cur = append(cur, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[len(cur):]
		out = append(out, cur)
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

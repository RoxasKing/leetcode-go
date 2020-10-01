package main

/*
  给定一个二叉树，返回它的 前序 遍历。

  进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

// DFS + Recursion
func preorderTraversal(root *TreeNode) []int {
	var out []int
	dfs(root, &out)
	return out
}

func dfs(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}
	*out = append(*out, root.Val)
	dfs(root.Left, out)
	dfs(root.Right, out)
}

// Stack + Iteration
func preorderTraversal2(root *TreeNode) []int {
	stack := []*TreeNode{}
	node := root
	var out []int
	for len(stack) != 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			out = append(out, node.Val)
			node = node.Left
		}
		last := len(stack) - 1
		if stack[last].Right != nil {
			node = stack[last].Right
		}
		stack = stack[:last]
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

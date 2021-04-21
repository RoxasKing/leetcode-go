package main

/*
  给定一个二叉树，返回它的 后序 遍历。

  进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

// DFS + Recursion
func postorderTraversal(root *TreeNode) []int {
	var out []int
	dfs(root, &out)
	return out
}

func dfs(node *TreeNode, out *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, out)
	dfs(node.Right, out)
	*out = append(*out, node.Val)
}

// Stack + Iteration
func postorderTraversal2(root *TreeNode) []int {
	var out []int
	var stack []*TreeNode
	var prev *TreeNode
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		top := stack[len(stack)-1]
		if top.Right != nil && top.Right != prev {
			node = top.Right
		} else {
			stack = stack[:len(stack)-1]
			out = append(out, top.Val)
			prev = top
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

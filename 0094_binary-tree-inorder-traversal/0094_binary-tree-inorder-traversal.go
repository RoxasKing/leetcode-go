package leetcode

/*
  给定一个二叉树，返回它的中序 遍历。
*/

// Recursion
func inorderTraversal(root *TreeNode) []int {
	var out []int
	var recursion func(*TreeNode)
	recursion = func(node *TreeNode) {
		if node == nil {
			return
		}
		recursion(node.Left)
		out = append(out, node.Val)
		recursion(node.Right)
	}
	recursion(root)
	return out
}

// Stack
func inorderTraversal2(root *TreeNode) []int {
	var out []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			top := stack[len(stack)-1]
			out = append(out, top.Val)
			root = top.Right
			stack = stack[:len(stack)-1]
		}
	}
	return out
}

// Morris Traversal
func inorderTraversal3(root *TreeNode) []int {
	var out []int
	for root != nil {
		if root.Left != nil {
			pre := root.Left
			for pre.Right != nil && pre.Right != root {
				pre = pre.Right
			}
			if pre.Right != root {
				pre.Right = root
				root = root.Left
				continue
			}
			pre.Right = nil
		}
		out = append(out, root.Val)
		root = root.Right
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

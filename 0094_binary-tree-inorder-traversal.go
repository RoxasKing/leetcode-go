package leetcode

/*
  给定一个二叉树，返回它的中序 遍历。
*/

// Recursive
func inorderTraversal(root *TreeNode) []int {
	var out []int
	var traversal func(*TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		out = append(out, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return out
}

// Stack
func inorderTraversal2(root *TreeNode) []int {
	var out []int
	var stack []*TreeNode
	for node := root; node != nil || len(stack) != 0; {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) != 0 {
			out = append(out, stack[len(stack)-1].Val)
			node = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return out
}

// Morris Traversal
func inorderTraversal3(root *TreeNode) []int {
	var out []int
	for cur := root; cur != nil; {
		if cur.Left != nil {
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right != cur {
				pre.Right, cur = cur, cur.Left
				continue
			}
			pre.Right = nil
			out = append(out, cur.Val)
			cur = cur.Right
		} else {
			out = append(out, cur.Val)
			cur = cur.Right
		}
	}
	return out
}

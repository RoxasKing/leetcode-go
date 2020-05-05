package leetcode

/*
  给定一个二叉树，判断其是否是一个有效的二叉搜索树。
  假设一个二叉搜索树具有如下特征：
    节点的左子树只包含小于当前节点的数。
    节点的右子树只包含大于当前节点的数。
    所有左子树和右子树自身必须也是二叉搜索树。
*/

// DFS
func isValidBST(root *TreeNode) bool {
	return dfsIsValidBST(-1<<63, 1<<63-1, root)
}

func dfsIsValidBST(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return dfsIsValidBST(min, root.Val, root.Left) && dfsIsValidBST(root.Val, max, root.Right)
}

// Stack
func isValidBST2(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := -1 << 63
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

// Morris Traversal
func isValidBST3(root *TreeNode) bool {
	inorder := -1 << 63
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
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

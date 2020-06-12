package codinginterviews

/*
  给定一棵二叉搜索树，请找出其中第k大的节点。

  限制：
    1 ≤ k ≤ 二叉搜索树元素个数
*/

// Morris Traversal
func kthLargest(root *TreeNode, k int) int {
	var index int
	for root != nil {
		if root.Right != nil {
			pre := root.Right
			for pre.Left != nil && pre.Left != root {
				pre = pre.Left
			}
			if pre.Left != root {
				pre.Left = root
				root = root.Right
				continue
			}
			index++
			if index == k {
				return root.Val
			}
			root = root.Left
			pre.Left = nil
		} else {
			index++
			if index == k {
				return root.Val
			}
			root = root.Left
		}
	}
	return -1
}

// DFS
func kthLargest2(root *TreeNode, k int) int {
	var out int
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if ok := dfs(root.Right); ok {
			return true
		}
		if k == 1 {
			out = root.Val
			return true
		}
		k--
		if ok := dfs(root.Left); ok {
			return true
		}
		return false
	}
	dfs(root)
	return out
}

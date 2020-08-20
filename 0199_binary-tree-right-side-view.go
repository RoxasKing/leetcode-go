package leetcode

/*
  给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/

// DFS + Recursion
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var out []int
	dfsRightSideView(root, &out, 0)
	return out
}

func dfsRightSideView(node *TreeNode, out *[]int, deep int) {
	if len((*out)) == deep {
		*out = append(*out, node.Val)
	}
	deep++
	if node.Right != nil {
		dfsRightSideView(node.Right, out, deep)
	}
	if node.Left != nil {
		dfsRightSideView(node.Left, out, deep)
	}
}

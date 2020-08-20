package main

/*
  给定一个二叉树，判断其是否是一个有效的二叉搜索树。

  假设一个二叉搜索树具有如下特征：
    节点的左子树只包含小于当前节点的数。
    节点的右子树只包含大于当前节点的数。
    所有左子树和右子树自身必须也是二叉搜索树。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/validate-binary-search-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Recursion
func isValidBST(root *TreeNode) bool {
	return dfs(-1<<63, 1<<63-1, root)
}

func dfs(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	} else if root.Val <= min || root.Val >= max {
		return false
	}
	return dfs(min, root.Val, root.Left) && dfs(root.Val, max, root.Right)
}

// DFS + Stack
func isValidBST2(root *TreeNode) bool {
	var stack []*TreeNode
	cur := root
	preVal := -1 << 63
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if preVal >= cur.Val {
			return false
		}
		preVal, cur = cur.Val, cur.Right
	}
	return true
}

// Morris Traversal
func isValidBST3(root *TreeNode) bool {
	cur := root
	preVal := -1 << 31
	for cur != nil {
		if cur.Left != nil {
			curPre := cur.Left
			for curPre.Right != nil && curPre.Right != cur {
				curPre = curPre.Right
			}
			if curPre.Right != cur {
				curPre.Right, cur = cur, cur.Left
				continue
			}
			curPre.Right = nil
		}
		if preVal >= cur.Val {
			return false
		}
		preVal, cur = cur.Val, cur.Right
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

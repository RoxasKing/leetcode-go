package main

/*
  Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

  Example 1:
    Input: root = [1,2,3,null,5,null,4]
    Output: [1,3,4]

  Example 2:
    Input: root = [1,null,3]
    Output: [1,3]

  Example 3:
    Input: root = []
    Output: []

  Constraints:
    1. The number of nodes in the tree is in the range [0, 100].
    2. -100 <= Node.val <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Recursion

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	out := []int{}
	dfs(root, 0, &out)
	return out
}

func dfs(root *TreeNode, depth int, out *[]int) {
	if root == nil {
		return
	}
	if depth > len(*out)-1 {
		*out = append(*out, root.Val)
	} else {
		(*out)[depth] = root.Val
	}
	dfs(root.Left, depth+1, out)
	dfs(root.Right, depth+1, out)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

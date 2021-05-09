package main

/*
  Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where each path's sum equals targetSum.

  A leaf is a node with no children.

  Example 1:
    Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
    Output: [[5,4,11,2],[5,8,4,5]]

  Example 2:
    Input: root = [1,2,3], targetSum = 5
    Output: []

  Example 3:
    Input: root = [1,2], targetSum = 0
    Output: []

  Constraints:
    1. The number of nodes in the tree is in the range [0, 5000].
    2. -1000 <= Node.val <= 1000
    3. -1000 <= targetSum <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/path-sum-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	out := [][]int{}
	dfs(root, targetSum-root.Val, []int{root.Val}, &out)
	return out
}

func dfs(root *TreeNode, targetSum int, cur []int, out *[][]int) {
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			*out = append(*out, tmp)
		}
		return
	}

	if root.Left != nil {
		targetSum -= root.Left.Val
		cur = append(cur, root.Left.Val)
		dfs(root.Left, targetSum, cur, out)
		cur = cur[:len(cur)-1]
		targetSum += root.Left.Val
	}

	if root.Right != nil {
		targetSum -= root.Right.Val
		cur = append(cur, root.Right.Val)
		dfs(root.Right, targetSum, cur, out)
		// cur = cur[:len(cur)-1]
		// targetSum += root.Right.Val
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

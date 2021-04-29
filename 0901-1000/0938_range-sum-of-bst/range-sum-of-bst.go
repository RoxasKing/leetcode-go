package main

/*
  Given the root node of a binary search tree, return the sum of values of all nodes with a value in the range [low, high].

  Example 1:
    Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
    Output: 32

  Example 2:
    Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
    Output: 23

  Constraints:
    1. The number of nodes in the tree is in the range [1, 2 * 10^4].
    2. 1 <= Node.val <= 10^5
    3. 1 <= low <= high <= 10^5
    4. All Node.val are unique.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/range-sum-of-bst
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	out := 0
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
		if low <= root.Val && root.Val <= high {
			out += root.Val
		}
		root = root.Right
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

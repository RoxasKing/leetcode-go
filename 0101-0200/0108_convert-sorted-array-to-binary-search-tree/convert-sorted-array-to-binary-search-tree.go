package main

/*
  Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.

  A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

  Example 1:
    Input: nums = [-10,-3,0,5,9]
    Output: [0,-3,9,-10,null,5]
    Explanation: [0,-10,5,null,-3,null,9] is also accepted:

  Example 2:
    Input: nums = [1,3]
    Output: [3,1]
    Explanation: [1,3] and [3,1] are both a height-balanced BSTs.

  Constraints:
    1. 1 <= nums.length <= 10^4
    2. -10^4 <= nums[i] <= 10^4
    3. nums is sorted in a strictly increasing order.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	m := n >> 1
	root := &TreeNode{Val: nums[m]}
	root.Left = sortedArrayToBST(nums[:m])
	root.Right = sortedArrayToBST(nums[m+1:])
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

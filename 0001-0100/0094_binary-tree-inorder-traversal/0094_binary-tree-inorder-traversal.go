package main

/*
  Given the root of a binary tree, return the inorder traversal of its nodes' values.

  Example 1:
    Input: root = [1,null,2,3]
    Output: [1,3,2]

  Example 2:
    Input: root = []
    Output: []

  Example 3:
    Input: root = [1]
    Output: [1]

  Example 4:
    Input: root = [1,2]
    Output: [2,1]

  Example 5:
    Input: root = [1,null,2]
    Output: [1,2]

  Constraints:
    1. The number of nodes in the tree is in the range [0, 100].
    2. -100 <= Node.val <= 100

  Follow up:
    Recursive solution is trivial, could you do it iteratively?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Morris Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	out := []int{}
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

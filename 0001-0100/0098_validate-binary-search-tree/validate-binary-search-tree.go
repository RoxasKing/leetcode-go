package main

/*
  Given the root of a binary tree, determine if it is a valid binary search tree (BST).

  A valid BST is defined as follows:
    1. The left subtree of a node contains only nodes with keys less than the node's key.
    2. The right subtree of a node contains only nodes with keys greater than the node's key.
    3. Both the left and right subtrees must also be binary search trees.

  Example 1:
    Input: root = [2,1,3]
    Output: true

  Example 2:
    Input: root = [5,1,4,null,null,3,6]
    Output: false
    Explanation: The root node's value is 5 but its right child's value is 4.

  Constraints:
    1. The number of nodes in the tree is in the range [1, 10^4].
    2. -2^31 <= Node.val <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/validate-binary-search-tree
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
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	for root != nil {
		if root.Left != nil {
			rootPre := root.Left
			for rootPre.Right != nil && rootPre.Right != root {
				rootPre = rootPre.Right
			}
			if rootPre.Right != root {
				rootPre.Right = root
				root = root.Left
				continue
			}
			rootPre.Right = nil
		}
		if prev != nil && prev.Val >= root.Val {
			return false
		}
		prev, root = root, root.Right
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

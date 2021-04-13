package main

/*
  You are given the root of a binary search tree (BST), where exactly two nodes of the tree were swapped by mistake. Recover the tree without changing its structure.

  Follow up: A solution using O(n) space is pretty straight forward. Could you devise a constant space solution?

  Example 1:
    Input: root = [1,3,null,null,2]
    Output: [3,1,null,null,2]
    Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST valid.

  Example 2:
    Input: root = [3,1,4,null,null,2]
    Output: [2,1,4,null,null,3]
    Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes the BST valid.

  Constraints:
    1. The number of nodes in the tree is in the range [2, 1000].
    2. -2^31 <= Node.val <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/recover-binary-search-tree
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

// Morris Traversal
func recoverTree(root *TreeNode) {
	var last, a, b *TreeNode
	node := root
	for node != nil {
		if node.Left != nil {
			prev := node.Left
			for prev.Right != nil && prev.Right != node {
				prev = prev.Right
			}
			if prev.Right != node {
				prev.Right = node
				node = node.Left
				continue
			}
			prev.Right = nil
		}
		if last != nil && node.Val < last.Val {
			if a == nil {
				a = last
			}
			b = node
		}
		last = node
		node = node.Right
	}
	a.Val, b.Val = b.Val, a.Val
}

// Stack
func recoverTree2(root *TreeNode) {
	var last, a, b *TreeNode
	node := root
	stack := []*TreeNode{}
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last != nil && node.Val < last.Val {
			if a == nil {
				a = last
			}
			b = node
		}
		last = node
		node = node.Right
	}
	a.Val, b.Val = b.Val, a.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

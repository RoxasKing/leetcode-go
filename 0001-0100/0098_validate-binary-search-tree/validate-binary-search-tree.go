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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
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
	stack := []*TreeNode{}
	node := root
	val := -1 << 63
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val <= val {
			return false
		}
		val = node.Val
		node = node.Right
	}
	return true
}

// Morris Traversal
func isValidBST3(root *TreeNode) bool {
	node := root
	val := -1 << 63
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
		if val >= node.Val {
			return false
		}
		val = node.Val
		node = node.Right
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

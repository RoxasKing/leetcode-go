package main

/*
  Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

  Example 1:
    Input: root = [3,9,20,null,null,15,7]
    Output: [[3],[9,20],[15,7]]

  Example 2:
    Input: root = [1]
    Output: [[1]]

  Example 3:
    Input: root = []
    Output: []

  Constraints:
    1. The number of nodes in the tree is in the range [0, 2000].
    2. -1000 <= Node.val <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	out := [][]int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		tmp := []int{}
		for _, t := range q {
			tmp = append(tmp, t.Val)
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		out = append(out, tmp)
		q = q[size:]
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

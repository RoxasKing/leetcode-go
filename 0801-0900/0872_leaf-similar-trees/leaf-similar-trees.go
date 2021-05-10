package main

/*
  Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.

  For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

  Two binary trees are considered leaf-similar if their leaf value sequence is the same.

  Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

  Example 1:
    Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
    Output: true

  Example 2:
    Input: root1 = [1], root2 = [1]
    Output: true

  Example 3:
    Input: root1 = [1], root2 = [2]
    Output: false

  Example 4:
    Input: root1 = [1,2], root2 = [2,2]
    Output: true

  Example 5:
    Input: root1 = [1,2,3], root2 = [1,3,2]
    Output: false

  Constraints:
    1. The number of nodes in each tree will be in the range [1, 200].
    2. Both of the given trees will have values in the range [0, 200].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/leaf-similar-trees
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack + DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := getAllLeaf(root1)
	arr2 := getAllLeaf(root2)
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func getAllLeaf(root *TreeNode) []int {
	out := []int{}
	stk := []*TreeNode{root}
	for len(stk) > 0 {
		top := len(stk) - 1
		root, stk = stk[top], stk[:top]
		if root.Left == nil && root.Right == nil {
			out = append(out, root.Val)
		}
		if root.Right != nil {
			stk = append(stk, root.Right)
		}
		if root.Left != nil {
			stk = append(stk, root.Left)
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

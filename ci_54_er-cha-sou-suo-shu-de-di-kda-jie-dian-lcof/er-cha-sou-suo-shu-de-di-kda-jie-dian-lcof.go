package main

/*
  给定一棵二叉搜索树，请找出其中第k大的节点。

  示例 1:
    输入: root = [3,1,4,null,2], k = 1
       3
      / \
     1   4
      \
       2
    输出: 4

  示例 2:
    输入: root = [5,3,6,2,4,null,null,1], k = 3
           5
          / \
         3   6
        / \
       2   4
      /
     1
    输出: 4

  限制：
    1 ≤ k ≤ 二叉搜索树元素个数

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof
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
func kthLargest(root *TreeNode, k int) int {
	out := root.Val
	i := 0
	n := root
	for n != nil {
		if n.Right != nil {
			nNext := n.Right
			for nNext.Left != nil && nNext.Left != n {
				nNext = nNext.Left
			}
			if nNext.Left != n {
				nNext.Left = n
				n = n.Right
				continue
			}
			nNext.Left = nil
		}
		i++
		out = n.Val
		if i == k {
			break
		}
		n = n.Left
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func kthLargest2(root *TreeNode, k int) int {
	var out int
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if ok := dfs(root.Right); ok {
			return true
		}
		if k == 1 {
			out = root.Val
			return true
		}
		k--
		if ok := dfs(root.Left); ok {
			return true
		}
		return false
	}
	dfs(root)
	return out
}

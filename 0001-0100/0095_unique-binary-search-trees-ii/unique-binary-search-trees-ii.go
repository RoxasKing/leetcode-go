package main

/*
  Given an integer n, return all the structurally unique BST's (binary search trees), which has exactly n nodes of unique values from 1 to n. Return the answer in any order.

  Example 1:
    Input: n = 3
    Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

  Example 2:
    Input: n = 1
    Output: [[1]]

  Constraints:
    1 <= n <= 8

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
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
func generateTrees(n int) []*TreeNode {
	return dfs(1, n)
}

func dfs(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	out := []*TreeNode{}
	for i := l; i <= r; i++ {
		ls := dfs(l, i-1)
		rs := dfs(i+1, r)
		for _, l := range ls {
			for _, r := range rs {
				out = append(out, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

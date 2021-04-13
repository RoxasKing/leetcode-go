package main

/*
  小扣有一个根结点为 root 的二叉树模型，初始所有结点均为白色，可以用蓝色染料给模型结点染色，模型的每个结点有一个 val 价值。小扣出于美观考虑，希望最后二叉树上每个蓝色相连部分的结点个数不能超过 k 个，求所有染成蓝色的结点价值总和最大是多少？

  示例 1：
    输入：root = [5,2,3,4], k = 2
    输出：12
    解释：结点 5、3、4 染成蓝色，获得最大的价值 5+3+4=12

  示例 2：
    输入：root = [4,1,3,9,null,null,2], k = 2
    输出：16
    解释：结点 4、3、9 染成蓝色，获得最大的价值 4+3+9=16

  提示：
    1. 1 <= k <= 10
    2. 1 <= val <= 10000
    3. 1 <= 结点数量 <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/er-cha-shu-ran-se-UGC
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Dynamic Programming

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxValue(root *TreeNode, k int) int {
	dp := make(map[*TreeNode][]int)
	return dfs(dp, root, k, k)
}

func dfs(dp map[*TreeNode][]int, root *TreeNode, remain, k int) int {
	if root == nil {
		return 0
	}

	if dp[root] == nil {
		dp[root] = make([]int, k+1)
	}

	if dp[root][remain] > 0 {
		return dp[root][remain]
	}

	out := dfs(dp, root.Left, k, k) + dfs(dp, root.Right, k, k) // root 不染色

	for i := 0; i <= remain-1; i++ {
		l1 := dfs(dp, root.Left, i, k)           // left 染色 i
		r1 := dfs(dp, root.Right, remain-1-i, k) // right 染色 remain-i
		out = Max(out, root.Val+l1+r1)           // root 染色
	}
	dp[root][remain] = out
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

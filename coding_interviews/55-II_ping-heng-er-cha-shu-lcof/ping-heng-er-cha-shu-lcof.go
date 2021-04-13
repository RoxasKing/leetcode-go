package main

/*
  输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

  示例 1:
    给定二叉树 [3,9,20,null,null,15,7]

        3
       / \
      9  20
        /  \
       15   7
    返回 true 。

  示例 2:
    给定二叉树 [1,2,2,3,3,null,null,4,4]

           1
          / \
         2   2
        / \
       3   3
      / \
     4   4
    返回 false 。

  限制：
    0 <= 树的结点个数 <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof
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
func isBalanced(root *TreeNode) bool {
	_, ok := dfs(root)
	return ok
}

func dfs(root *TreeNode) (depth int, ok bool) {
	if root == nil {
		return 0, true
	}
	dL, okL := dfs(root.Left)
	if !okL {
		return 0, false
	}
	dR, okR := dfs(root.Right)
	if !okR {
		return 0, false
	}
	return Max(dL, dR) + 1, Abs(dL-dR) < 2
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

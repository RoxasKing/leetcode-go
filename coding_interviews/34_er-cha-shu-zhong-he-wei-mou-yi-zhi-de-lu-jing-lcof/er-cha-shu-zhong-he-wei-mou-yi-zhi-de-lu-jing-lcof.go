package main

/*
  输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

  示例:
    给定如下二叉树，以及目标和 sum = 22，

                  5
                 / \
                4   8
               /   / \
              11  13  4
             /  \    / \
            7    2  5   1
    返回:
    [
       [5,4,11,2],
       [5,8,4,5]
    ]

  提示：
    节点总数 <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Backtrack
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	out := [][]int{}
	dfs(root, sum, root.Val, []int{root.Val}, &out)
	return out
}

func dfs(node *TreeNode, sum, cur int, arr []int, out *[][]int) {
	if node.Left == nil && node.Right == nil {
		if cur == sum {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			*out = append(*out, tmp)
		}
		return
	}
	if node.Left != nil {
		cur += node.Left.Val
		arr = append(arr, node.Left.Val)
		dfs(node.Left, sum, cur, arr, out)
		arr = arr[:len(arr)-1]
		cur -= node.Left.Val
	}
	if node.Right != nil {
		cur += node.Right.Val
		arr = append(arr, node.Right.Val)
		dfs(node.Right, sum, cur, arr, out)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

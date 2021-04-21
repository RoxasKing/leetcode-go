package main

/*
  给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

  提示：
    树中至少有 2 个节点。
    本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getMinimumDifference(root *TreeNode) int {
	node := root
	min := 1<<31 - 1
	var pre *TreeNode
	for node != nil {
		if node.Left != nil {
			pre := node.Left
			for pre.Right != nil && pre.Right != node {
				pre = pre.Right
			}
			if pre.Right != node {
				pre.Right = node
				node = node.Left
				continue
			}
			pre.Right = nil
		}
		if pre != nil {
			min = Min(min, node.Val-pre.Val)
		}
		pre = node
		node = node.Right
	}
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

package main

/*
  给定一个二叉树，我们在树的节点上安装摄像头。
  节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
  计算监控树的所有节点所需的最小摄像头数量。

  提示：
    给定树的节点数的范围是 [1, 1000]。
    每个节点的值都是 0。
*/

// DFS + Recursion
func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	} else if isLeaf(root) {
		return 1
	}
	out, leafHaveCamera, leafLeafHaveCamera := dfs(root)
	if !(leafHaveCamera || leafLeafHaveCamera) {
		out++
	}
	return out
}

func dfs(node *TreeNode) (int, bool, bool) {
	if isLeaf(node) {
		return 0, false, false
	}
	var lCount, rCount int
	var lChildHaveCamera, lChildChildHaveCamera, rChildHaveCamera, rChildChildHaveCamera bool
	if node.Left != nil {
		lCount, lChildHaveCamera, lChildChildHaveCamera = dfs(node.Left)
	}
	if node.Right != nil {
		rCount, rChildHaveCamera, rChildChildHaveCamera = dfs(node.Right)
	}
	count := lCount + rCount
	// 若当前节点有一个子节点是叶子节点
	// 或当前节点的非空子节点的子节点，以及其子节点的子节点没有摄像头
	// 则当前节点必须有摄像头
	curNodeHaveCamera :=
		node.Left != nil && (isLeaf(node.Left) || !lChildHaveCamera && !lChildChildHaveCamera) ||
			node.Right != nil && (isLeaf(node.Right) || !rChildHaveCamera && !rChildChildHaveCamera)
	if curNodeHaveCamera {
		count++
	}
	return count, curNodeHaveCamera, lChildHaveCamera || rChildHaveCamera
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

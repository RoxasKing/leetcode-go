package leetcode

import "math/rand"

/*
  将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
  本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
  示例:
    给定有序数组: [-10,-3,0,5,9],
    一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
          0
         / \
       -3   9
       /   /
     -10  5
*/

func sortedArrayToBST(nums []int) *TreeNode {
	var buildBST func(int, int) *TreeNode
	buildBST = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := l + (r-l)>>1
		root := &TreeNode{Val: nums[m]}
		root.Left = buildBST(l, m-1)
		root.Right = buildBST(m+1, r)
		return root
	}
	return buildBST(0, len(nums)-1)
}

func sortedArrayToBST2(nums []int) *TreeNode {
	var buildBST func(int, int) *TreeNode
	buildBST = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := l + (r+1-l)>>1
		root := &TreeNode{Val: nums[m]}
		root.Left = buildBST(l, m-1)
		root.Right = buildBST(m+1, r)
		return root
	}
	return buildBST(0, len(nums)-1)
}

func sortedArrayToBST3(nums []int) *TreeNode {
	var buildBST func(int, int) *TreeNode
	buildBST = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := l + (r+rand.Intn(2)-l)>>1
		root := &TreeNode{Val: nums[m]}
		root.Left = buildBST(l, m-1)
		root.Right = buildBST(m+1, r)
		return root
	}
	return buildBST(0, len(nums)-1)
}

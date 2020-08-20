package main

/*
  将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

  本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := l + (r-l+1)>>1
	return &TreeNode{
		Val:   nums[m],
		Left:  buildBST(nums, l, m-1),
		Right: buildBST(nums, m+1, r),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

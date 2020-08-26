package main

/*
  给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

  说明：
  你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

  进阶：
    如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Morris Traversal
func kthSmallest(root *TreeNode, k int) int {
	var index int
	node := root
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
		index++
		if index == k {
			break
		}
		node = node.Right
	}
	return node.Val
}

// Stack
func kthSmallest2(root *TreeNode, k int) int {
	var index int
	var stack []*TreeNode
	node := root
	for len(stack) != 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) != 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			index++
			if index == k {
				break
			}
			node = node.Right
		}
	}
	return node.Val
}

// Recursion
func kthSmallest3(root *TreeNode, k int) int {
	var out int
	var index int
	var recursion func(*TreeNode) bool
	recursion = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if recursion(node.Left) {
			return true
		}
		index++
		if index == k {
			out = node.Val
			return true
		}
		return recursion(node.Right)

	}
	recursion(root)
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

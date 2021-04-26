package main

import "sort"

/*
  给出一个完全二叉树，求出该树的节点个数。

  说明：
    完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

  示例:
    输入:
        1
       / \
      2   3
     / \  /
    4  5 6
    输出: 6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-complete-tree-nodes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// Binary Search + Bit Manipulation
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	for n := root; n.Left != nil; n = n.Left {
		depth++
	}
	return sort.Search(1<<(depth+1), func(i int) bool {
		if i <= 1<<depth {
			return false
		}
		bits := 1 << (depth - 1)
		n := root
		for n != nil && bits > 0 {
			if bits&i == 0 {
				n = n.Left
			} else {
				n = n.Right
			}
			bits >>= 1
		}
		return n == nil
	}) - 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

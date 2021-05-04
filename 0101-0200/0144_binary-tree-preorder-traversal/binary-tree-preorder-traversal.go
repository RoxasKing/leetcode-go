package main

/*
  Given the root of a binary tree, return the preorder traversal of its nodes' values.

  Example 1:
    Input: root = [1,null,2,3]
    Output: [1,2,3]

  Example 2:
    Input: root = []
    Output: []

  Example 3:
    Input: root = [1]
    Output: [1]

  Example 4:
    Input: root = [1,2]
    Output: [1,2]

  Example 5:
    Input: root = [1,null,2]
    Output: [1,2]

  Constraints:
    1. The number of nodes in the tree is in the range [0, 100].
    2. -100 <= Node.val <= 100

  Follow up: Recursive solution is trivial, could you do it iteratively?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS
func preorderTraversal(root *TreeNode) []int {
	var out []int
	dfs(root, &out)
	return out
}

func dfs(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}
	*out = append(*out, root.Val)
	dfs(root.Left, out)
	dfs(root.Right, out)
}

// Stack
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := []int{}
	stk := []*TreeNode{root}
	for len(stk) > 0 {
		top := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		out = append(out, top.Val)
		if top.Right != nil {
			stk = append(stk, top.Right)
		}
		if top.Left != nil {
			stk = append(stk, top.Left)
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris Traversal
func preorderTraversal3(root *TreeNode) []int {
	out := []int{}
	node := root
	for node != nil {
		if node.Left != nil {
			pre := node.Left
			for pre.Right != nil && pre.Right != node {
				pre = pre.Right
			}
			if pre.Right != node {
				pre.Right = node
				out = append(out, node.Val)
				node = node.Left
				continue
			}
			pre.Right = nil
			node = node.Right
		} else {
			out = append(out, node.Val)
			node = node.Right
		}
	}
	return out
}

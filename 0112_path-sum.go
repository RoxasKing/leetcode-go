package main

/*
  给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
  说明: 叶子节点是指没有子节点的节点。
  示例:
  给定如下二叉树，以及目标和 sum = 22，
                5
               / \
              4   8
             /   / \
            11  13  4
           /  \      \
          7    2      1
  返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
*/

// Queue
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	type node struct {
		t *TreeNode
		s int
	}
	queue := []node{{root, root.Val}}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.t.Left == nil && cur.t.Right == nil && cur.s == sum {
			return true
		}
		if cur.t.Left != nil {
			queue = append(queue, node{t: cur.t.Left, s: cur.s + cur.t.Left.Val})
		}
		if cur.t.Right != nil {
			queue = append(queue, node{t: cur.t.Right, s: cur.s + cur.t.Right.Val})
		}
	}
	return false
}

// Recursion
func hasPathSum2(root *TreeNode, sum int) bool {
	var help func(*TreeNode, int) bool
	help = func(root *TreeNode, sum int) bool {
		if root == nil {
			return false
		}
		sum -= root.Val
		if root.Left == nil && root.Right == nil {
			return sum == 0
		}
		return help(root.Left, sum) || help(root.Right, sum)
	}
	return help(root, sum)
}

package main

/*
  给定一个二叉树，检查它是否是镜像对称的。
  例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
        1
       / \
      2   2
     / \ / \
    3  4 4  3
  但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
        1
       / \
      2   2
       \   \
       3    3
*/

// BFS + Recursion
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return bfs(root.Left, root.Right)
}

func bfs(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return bfs(l.Left, r.Right) && bfs(l.Right, r.Left)
}

// BFS + Iteration
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) != 0 {
		l, r := queue[0], queue[1]
		queue = queue[2:]
		if l == nil && r == nil {
		} else if l != nil || r != nil || l.Val != r.Val {
			return false
		} else {
			queue = append(queue, l.Left, r.Right, l.Right, r.Left)
		}
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

package leetcode

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

// Recursion
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var compare func(*TreeNode, *TreeNode) bool
	compare = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		} else if a == nil || b == nil {
			return false
		} else if a.Val != b.Val {
			return false
		}
		return compare(a.Left, b.Right) && compare(a.Right, b.Left)
	}
	return compare(root.Left, root.Right)
}

// Iteration
func isSymmetric2(root *TreeNode) bool {
	var stack []*TreeNode
	var l, r *TreeNode
	stack = append(stack, root, root)
	for len(stack) != 0 {
		l, stack = stack[0], stack[1:]
		r, stack = stack[0], stack[1:]
		if l == nil && r == nil {
			continue
		} else if l == nil || r == nil {
			return false
		} else if l.Val != r.Val {
			return false
		}
		stack = append(stack, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}

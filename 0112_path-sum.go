package leetcode

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

// Stack
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	type elem struct {
		node *TreeNode
		sum  int
	}
	stack := []*elem{{root, root.Val}}
	for len(stack) != 0 {
		index := len(stack)
		for i := range stack {
			if stack[i].node.Left == nil && stack[i].node.Right == nil && stack[i].sum == sum {
				return true
			}
			if stack[i].node.Left != nil {
				stack = append(
					stack,
					&elem{stack[i].node.Left, stack[i].sum + stack[i].node.Left.Val},
				)
			}
			if stack[i].node.Right != nil {
				stack = append(
					stack,
					&elem{stack[i].node.Right, stack[i].sum + stack[i].node.Right.Val},
				)
			}
		}
		stack = stack[index:]
	}
	return false
}

// Recursive
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum2(root.Left, sum) || hasPathSum2(root.Right, sum)
}

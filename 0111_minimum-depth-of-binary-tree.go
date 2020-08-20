package leetcode

/*
  给定一个二叉树，找出其最小深度。
  最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
  说明: 叶子节点是指没有子节点的节点。
  示例:
  给定二叉树 [3,9,20,null,null,15,7],
        3
       / \
      9  20
        /  \
       15   7
  返回它的最小深度  2.
*/

// Recursion
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	return Min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// Stack : width first
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type stack struct {
		node  *TreeNode
		depth int
	}
	stacks := []*stack{{root, 1}}
	for len(stacks) != 0 {
		if stacks[0].node.Left == nil && stacks[0].node.Right == nil {
			break
		}
		if stacks[0].node.Left != nil {
			stacks = append(stacks, &stack{stacks[0].node.Left, stacks[0].depth + 1})
		}
		if stacks[0].node.Right != nil {
			stacks = append(stacks, &stack{stacks[0].node.Right, stacks[0].depth + 1})
		}
		stacks = stacks[1:]
	}
	return stacks[0].depth
}

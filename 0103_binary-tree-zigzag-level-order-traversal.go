package leetcode

/*
  给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
  例如：
  给定二叉树 [3,9,20,null,null,15,7],
        3
       / \
      9  20
        /  \
       15   7
  返回锯齿形层次遍历如下：
    [
      [3],
      [20,9],
      [15,7]
    ]
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		out   [][]int
		isRev bool
		stack = []*TreeNode{root}
	)
	addChildNode := func(node *TreeNode) {
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	for len(stack) != 0 {
		var tmp []int
		curSize := len(stack)
		if !isRev {
			for i := range stack {
				tmp = append(tmp, stack[i].Val)
				addChildNode(stack[i])
			}
			isRev = true
		} else {
			for i := len(stack) - 1; i >= 0; i-- {
				tmp = append(tmp, stack[i].Val)
			}
			for i := range stack {
				addChildNode(stack[i])
			}
			isRev = false
		}
		out = append(out, tmp)
		stack = stack[curSize:]
	}
	return out
}

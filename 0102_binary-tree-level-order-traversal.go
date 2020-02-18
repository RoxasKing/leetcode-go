package leetcode

/*
  给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
  例如:
  给定二叉树: [3,9,20,null,null,15,7],
        3
       / \
      9  20
        /  \
       15   7
  返回其层次遍历结果：
    [
      [3],
      [9,20],
      [15,7]
    ]
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var out [][]int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		var tmp []int
		curSize := len(stack)
		for i := range stack {
			tmp = append(tmp, stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		out = append(out, tmp)
		stack = stack[curSize:]
	}
	return out
}

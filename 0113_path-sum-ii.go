package leetcode

/*
  给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
  说明: 叶子节点是指没有子节点的节点。
  示例:
  给定如下二叉树，以及目标和 sum = 22，
                5
               / \
              4   8
             /   / \
            11  13  4
           /  \    / \
          7    2  5   1
  返回:
    [
       [5,4,11,2],
       [5,8,4,5]
    ]
*/

// Recursion
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var out [][]int
	cur := []int{root.Val}
	var helper func(*TreeNode, int)
	helper = func(root *TreeNode, sum int) {
		if root.Left == nil && root.Right == nil {
			if sum == 0 {
				out = append(out, append(make([]int, 0, len(cur)), cur...))
			}
			return
		}
		if root.Left != nil {
			cur = append(cur, root.Left.Val)
			helper(root.Left, sum-root.Left.Val)
			cur = cur[:len(cur)-1]
		}
		if root.Right != nil {
			cur = append(cur, root.Right.Val)
			helper(root.Right, sum-root.Right.Val)
			cur = cur[:len(cur)-1]
		}
	}
	helper(root, sum-root.Val)
	return out
}

// Stack
func pathSum2(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	type elem struct {
		node *TreeNode
		sum  int
		arr  []int
	}
	var out [][]int
	stack := []*elem{{root, sum - root.Val, []int{root.Val}}}
	for len(stack) != 0 {
		index := len(stack)
		for i := range stack {
			if stack[i].node.Left == nil && stack[i].node.Right == nil {
				if stack[i].sum == 0 {
					out = append(out, append([]int{}, stack[i].arr...))
				}
				continue
			}
			if stack[i].node.Left != nil {
				arr := make([]int, len(stack[i].arr)+1)
				copy(arr, stack[i].arr)
				arr[len(arr)-1] = stack[i].node.Left.Val
				stack = append(
					stack,
					&elem{
						stack[i].node.Left,
						stack[i].sum - stack[i].node.Left.Val,
						arr,
					},
				)
			}
			if stack[i].node.Right != nil {
				arr := make([]int, len(stack[i].arr)+1)
				copy(arr, stack[i].arr)
				arr[len(arr)-1] = stack[i].node.Right.Val
				stack = append(
					stack,
					&elem{
						stack[i].node.Right,
						stack[i].sum - stack[i].node.Right.Val,
						arr,
					},
				)
			}
		}
		stack = stack[index:]
	}
	return out
}

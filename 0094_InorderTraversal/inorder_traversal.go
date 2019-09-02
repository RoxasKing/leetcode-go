package main

import "fmt"

//Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)
	stack := make([]*TreeNode, 0)
	t := root
	for t != nil || len(stack) != 0 {
		for t != nil {
			// 根节点放入栈底，并迭代左子树依次推入栈底
			stack = append(stack, t)
			t = t.Left
		}
		if len(stack) != 0 {
			// 将栈顶的左子树的值加入到输出中
			out = append(out, stack[len(stack)-1].Val)
			// 根节点变为右子树
			t = stack[len(stack)-1].Right
			// 将栈顶元素退出
			stack = stack[:len(stack)-1]
		}
	}
	return out
}

func main() {
	root := &TreeNode{Val: 1}
	left1 := &TreeNode{Val: 2}
	root.Left = left1
	right1 := &TreeNode{Val: 3}
	root.Right = right1
	left2 := &TreeNode{Val: 4}
	left1.Left = left2
	right2 := &TreeNode{Val: 5}
	left1.Right = right2
	fmt.Println(inorderTraversal(root))
}

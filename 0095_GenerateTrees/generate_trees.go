package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate_trees(1, n)
}

func generate_trees(start, end int) []*TreeNode {
	out := make([]*TreeNode, 0)
	if start > end {
		out = append(out, nil)
		return out
	}

	// 遍历根节点
	for i := start; i <= end; i++ {
		// 所有可能的左子树
		leftTrees := generate_trees(start, i-1)
		// 所有可能的右子树
		rightTrees := generate_trees(i+1, end)

		// 遍历当前节点下，所有可能的结构，并加入到输出中
		for j := range leftTrees {
			for k := range rightTrees {
				current := &TreeNode{Val: i}
				current.Left = leftTrees[j]
				current.Right = rightTrees[k]
				out = append(out, current)
			}
		}
	}
	return out
}

func main() {
	n := 3
	treeNodes := generateTrees(n)
	fmt.Println(treeNodes[0])
}

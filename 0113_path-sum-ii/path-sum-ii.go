package main

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

// DFS + Backtracking
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var out [][]int
	dfs(root, sum, root.Val, []int{root.Val}, &out)
	return out
}

func dfs(node *TreeNode, sum, curSum int, list []int, lists *[][]int) {
	if node.Left == nil && node.Right == nil {
		if curSum == sum {
			tmp := make([]int, len(list))
			copy(tmp, list)
			*lists = append(*lists, tmp)
		}
	}
	if node.Left != nil {
		dfs(node.Left, sum, curSum+node.Left.Val, append(list, node.Left.Val), lists)
	}
	if node.Right != nil {
		dfs(node.Right, sum, curSum+node.Right.Val, append(list, node.Right.Val), lists)
	}
}

// Stack
func pathSum2(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var out [][]int
	q := []*Node{{node: root, list: []int{root.Val}, sum: root.Val}}
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		if n.node.Left == nil && n.node.Right == nil && n.sum == sum {
			out = append(out, n.list)
		}
		if n.node.Left != nil {
			tmp := make([]int, len(n.list)+1)
			copy(tmp, n.list)
			tmp[len(n.list)] = n.node.Left.Val
			q = append(q, &Node{node: n.node.Left, list: tmp, sum: n.sum + n.node.Left.Val})
		}
		if n.node.Right != nil {
			tmp := make([]int, len(n.list)+1)
			copy(tmp, n.list)
			tmp[len(n.list)] = n.node.Right.Val
			q = append(q, &Node{node: n.node.Right, list: tmp, sum: n.sum + n.node.Right.Val})
		}
	}
	return out
}

type Node struct {
	node *TreeNode
	list []int
	sum  int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

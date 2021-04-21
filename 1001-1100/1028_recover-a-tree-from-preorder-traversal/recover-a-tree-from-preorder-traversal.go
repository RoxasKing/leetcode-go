package main

/*
  We run a preorder depth-first search (DFS) on the root of a binary tree.

  At each node in this traversal, we output D dashes (where D is the depth of this node), then we output the value of this node.  If the depth of a node is D, the depth of its immediate child is D + 1.  The depth of the root node is 0.

  If a node has only one child, that child is guaranteed to be the left child.

  Given the output S of this traversal, recover the tree and return its root.

  Example 1:
    Input: S = "1-2--3--4-5--6--7"
    Output: [1,2,5,3,4,6,7]

  Example 2:
    Input: S = "1-2--3---4-5--6---7"
    Output: [1,2,5,3,null,6,null,4,null,7]

  Example 3:
    Input: S = "1-401--349---90--88"
    Output: [1,401,null,349,88,90]

  Constraints:
    1. The number of nodes in the original tree is in the range [1, 1000].
    2. 1 <= Node.val <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {
	i := 0
	return dfs(S, &i, 0)
}

func dfs(S string, i *int, depth int) *TreeNode {
	d := 0
	j := *i
	for j < len(S) && S[j] == '-' {
		d++
		j++
	}
	if d != depth {
		return nil
	}
	*i = j
	val := 0
	for *i < len(S) && '0' <= S[*i] && S[*i] <= '9' {
		val = val*10 + int(S[*i]-'0')
		*i++
	}
	return &TreeNode{
		Val:   val,
		Left:  dfs(S, i, depth+1),
		Right: dfs(S, i, depth+1),
	}
}

// DFS
func recoverFromPreorder2(S string) *TreeNode {
	q := []*TreeNode{}
	for S != "" {
		dep := 0
		for S[0] == '-' {
			dep++
			S = S[1:]
		}
		val := 0
		for S != "" && S[0] != '-' {
			val *= 10
			val += int(S[0] - '0')
			S = S[1:]
		}
		node := &TreeNode{Val: val}
		if dep < len(q) {
			q = q[:dep]
			q[len(q)-1].Right = node
		} else if len(q) > 0 {
			q[len(q)-1].Left = node
		}
		q = append(q, node)
	}
	return q[0]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

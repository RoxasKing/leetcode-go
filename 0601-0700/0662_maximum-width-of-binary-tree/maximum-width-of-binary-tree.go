package main

/*
  Given a binary tree, write a function to get the maximum width of the given tree. The maximum width of a tree is the maximum width among all levels.

  The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the null nodes between the end-nodes are also counted into the length calculation.

  It is guaranteed that the answer will in the range of 32-bit signed integer.

  Example 1:
    Input:
               1
             /   \
            3     2
           / \     \
          5   3     9
    Output: 4
    Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).

  Example 2:
    Input:
              1
             /
            3
           / \
          5   3
    Output: 2
    Explanation: The maximum width existing in the third level with the length 2 (5,3).

  Example 3:
    Input:
              1
             / \
            3   2
           /
          5
    Output: 2
    Explanation: The maximum width existing in the second level with the length 2 (3,2).

  Example 4:
    Input:
              1
             / \
            3   2
           /     \
          5       9
         /         \
        6           7
    Output: 8
    Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).

  Constraints:
    The given binary tree will have between 1 and 3000 nodes.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-width-of-binary-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	q := []*AnnotatedNode{{node: root, depth: 0, pos: 0}}
	out, curDepth, left := 0, 0, 0
	for len(q) > 0 {
		an := q[0]
		q = q[1:]
		node, depth, pos := an.node, an.depth, an.pos
		if node != nil {
			q = append(q, &AnnotatedNode{node: node.Left, depth: depth + 1, pos: pos << 1})
			q = append(q, &AnnotatedNode{node: node.Right, depth: depth + 1, pos: pos<<1 + 1})
			if curDepth != depth {
				curDepth = depth
				left = pos
			}
			out = Max(out, pos+1-left)
		}
	}
	return out
}

type AnnotatedNode struct {
	node  *TreeNode
	depth int
	pos   int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

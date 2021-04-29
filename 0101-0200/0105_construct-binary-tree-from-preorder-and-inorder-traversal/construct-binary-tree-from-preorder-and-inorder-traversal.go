package main

/*
  Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

  Example 1:
    Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
    Output: [3,9,20,null,null,15,7]

  Example 2:
    Input: preorder = [-1], inorder = [-1]
    Output: [-1]

  Constraints:
    1. 1 <= preorder.length <= 3000
    2. inorder.length == preorder.length
    3. -3000 <= preorder[i], inorder[i] <= 3000
    4. preorder and inorder consist of unique values.
    5. Each value of inorder also appears in preorder.
    6. preorder is guaranteed to be the preorder traversal of the tree.
    7. inorder is guaranteed to be the inorder traversal of the tree.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, num := range inorder {
		inorderMap[num] = i
	}
	preorderIdx := 0
	return dfs(preorder, inorderMap, &preorderIdx, 0, len(inorder)-1)
}

func dfs(preorder []int, inorderMap map[int]int, preorderIdx *int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	rootVal := preorder[*preorderIdx]
	*preorderIdx++
	m := inorderMap[rootVal]
	return &TreeNode{
		Val:   rootVal,
		Left:  dfs(preorder, inorderMap, preorderIdx, l, m-1),
		Right: dfs(preorder, inorderMap, preorderIdx, m+1, r),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

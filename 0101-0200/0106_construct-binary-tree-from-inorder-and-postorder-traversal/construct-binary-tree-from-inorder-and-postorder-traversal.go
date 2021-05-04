package main

/*
  Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

  Example 1:
    Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
    Output: [3,9,20,null,null,15,7]

  Example 2:
    Input: inorder = [-1], postorder = [-1]
    Output: [-1]

  Constraints:
    1. 1 <= inorder.length <= 3000
    2. postorder.length == inorder.length
    3. -3000 <= inorder[i], postorder[i] <= 3000
    4. inorder and postorder consist of unique values.
    5. Each value of postorder also appears in inorder.
    6. inorder is guaranteed to be the inorder traversal of the tree.
    7. postorder is guaranteed to be the postorder traversal of the tree.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}
	n := len(postorder)
	pIdx := n - 1
	return dfs(inorderMap, postorder, &pIdx, 0, n-1)
}

func dfs(inorderMap map[int]int, postorder []int, pIdx *int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	rootVal := postorder[*pIdx]
	*pIdx--
	mid := inorderMap[rootVal]
	return &TreeNode{
		Val:   rootVal,
		Right: dfs(inorderMap, postorder, pIdx, mid+1, r),
		Left:  dfs(inorderMap, postorder, pIdx, l, mid-1),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

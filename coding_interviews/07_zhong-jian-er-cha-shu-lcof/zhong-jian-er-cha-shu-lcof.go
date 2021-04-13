package main

/*
  输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

  例如，给出
    前序遍历 preorder = [3,9,20,15,7]
    中序遍历 inorder = [9,3,15,20,7]
    返回如下的二叉树：
        3
       / \
      9  20
        /  \
       15   7

  限制：
    0 <= 节点个数 <= 5000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Recursion

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
	for i, v := range inorder {
		inorderMap[v] = i
	}
	i, l, r := 0, 0, len(preorder)-1
	return dfs(preorder, inorderMap, &i, l, r)
}

func dfs(preorder []int, inorderMap map[int]int, i *int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	val := preorder[*i]
	m := inorderMap[val]
	*i++
	return &TreeNode{
		Val:   val,
		Left:  dfs(preorder, inorderMap, i, l, m-1),
		Right: dfs(preorder, inorderMap, i, m+1, r),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

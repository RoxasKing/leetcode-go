package leetcode

/*
  根据一棵树的前序遍历与中序遍历构造二叉树。
  注意:
  你可以假设树中没有重复的元素。
  例如，给出
    前序遍历 preorder = [3,9,20,15,7]
    中序遍历 inorder = [9,3,15,20,7]
  返回如下的二叉树：
        3
       / \
      9  20
        /  \
       15   7
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	dict := make(map[int]int, len(inorder))
	for i := range inorder {
		dict[inorder[i]] = i
	}
	var preIndex int
	var helper func(int, int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left == right {
			return nil
		}
		rootVal := preorder[preIndex]
		index := dict[rootVal]
		preIndex++
		return &TreeNode{
			Val:   rootVal,
			Left:  helper(left, index),
			Right: helper(index+1, right),
		}
	}
	return helper(0, len(inorder))
}

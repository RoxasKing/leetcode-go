package main

/*
  根据一棵树的中序遍历与后序遍历构造二叉树。
  注意:
  你可以假设树中没有重复的元素。
  例如，给出
    中序遍历 inorder = [9,3,15,20,7]
    后序遍历 postorder = [9,15,7,20,3]
  返回如下的二叉树：
        3
       / \
      9  20
        /  \
       15   7
*/

func buildTree2(inorder []int, postorder []int) *TreeNode {
	dict := make(map[int]int, len(inorder))
	for i := range inorder {
		dict[inorder[i]] = i
	}
	postIndex := len(postorder) - 1
	var helper func(int, int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left == right {
			return nil
		}
		rootVal := postorder[postIndex]
		index := dict[rootVal]
		postIndex--
		root := &TreeNode{Val: rootVal}
		root.Right = helper(index+1, right)
		root.Left = helper(left, index)
		return root
	}
	return helper(0, len(inorder))
}

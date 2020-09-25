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

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for index, val := range inorder {
		inorderMap[val] = index
	}
	postorderIndex := len(postorder) - 1
	return build(postorder, &postorderIndex, inorderMap, 0, len(inorder)-1)
}

func build(postorder []int, postorderIndex *int, inorderMap map[int]int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	rootVal := postorder[*postorderIndex]
	inorderIndex := inorderMap[rootVal]
	*postorderIndex--
	return &TreeNode{
		Val:   rootVal,
		Right: build(postorder, postorderIndex, inorderMap, inorderIndex+1, r),
		Left:  build(postorder, postorderIndex, inorderMap, l, inorderIndex-1),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

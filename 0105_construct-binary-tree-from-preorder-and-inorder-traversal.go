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
	inorderDict := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderDict[v] = i
	}
	var preOrderIndex int
	var build func(int, int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		rootVal := preorder[preOrderIndex]
		inorderIndex := inorderDict[rootVal]
		preOrderIndex++
		return &TreeNode{
			Val:   rootVal,
			Left:  build(l, inorderIndex-1),
			Right: build(inorderIndex+1, r),
		}
	}
	return build(0, len(preorder)-1)
}

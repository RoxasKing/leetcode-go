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
	var preOrderPtr int
	var helper func(int, int) *TreeNode
	helper = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		val := preorder[preOrderPtr]
		inorderIndex := inorderDict[val]
		preOrderPtr++
		return &TreeNode{
			Val:   val,
			Left:  helper(l, inorderIndex-1),
			Right: helper(inorderIndex+1, r),
		}
	}
	return helper(0, len(preorder)-1)
}

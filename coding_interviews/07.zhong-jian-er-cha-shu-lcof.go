package codinginterviews

/*
  输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

  限制：
  0 <= 节点个数 <= 5000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderMap[v] = i
	}
	var preorderIndex int
	var build func(int, int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		rootVal := preorder[preorderIndex]
		root := &TreeNode{Val: rootVal}
		inorderIndex := inorderMap[rootVal]
		preorderIndex++
		root.Left = build(l, inorderIndex-1)
		root.Right = build(inorderIndex+1, r)
		return root
	}
	return build(0, len(inorder)-1)
}

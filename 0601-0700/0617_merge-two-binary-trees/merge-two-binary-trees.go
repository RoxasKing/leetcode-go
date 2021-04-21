package main

/*
  给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
  你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

  注意: 合并必须从两个树的根节点开始。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/merge-two-binary-trees
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 != nil {
		return &TreeNode{
			Val:   t1.Val + t2.Val,
			Left:  mergeTrees(t1.Left, t2.Left),
			Right: mergeTrees(t1.Right, t2.Right),
		}
	} else if t1 != nil {
		return t1
	} else if t2 != nil {
		return t2
	}
	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

package main

/*
  给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/subtree-of-another-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}
	queue := []*TreeNode{s}
	for len(queue) != 0 {
		s, queue = queue[0], queue[1:]
		if isEqual(s, t) {
			return true
		}
		if s.Left != nil {
			queue = append(queue, s.Left)
		}
		if s.Right != nil {
			queue = append(queue, s.Right)
		}
	}
	return false
}

func isEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return isEqual(a.Left, b.Left) && isEqual(a.Right, b.Right)
}

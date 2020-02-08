package leetcode

/*
  给定两个二叉树，编写一个函数来检验它们是否相同。
  如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// In stack
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	check := func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		} else if p == nil || q == nil {
			return false
		} else if p.Val != q.Val {
			return false
		}
		return true
	}
	var (
		queueP []*TreeNode
		queueQ []*TreeNode
	)
	queueP = append(queueP, p)
	queueQ = append(queueQ, q)
	for len(queueP) != 0 {
		p, queueP = queueP[len(queueP)-1], queueP[:len(queueP)-1]
		q, queueQ = queueQ[len(queueQ)-1], queueQ[:len(queueQ)-1]
		if !check(p, q) {
			return false
		}
		if p != nil {
			if !check(p.Left, q.Left) {
				return false
			}
			if p.Left != nil {
				queueP = append(queueP, p.Left)
				queueQ = append(queueQ, q.Left)
			}
			if !check(p.Right, q.Right) {
				return false
			}
			if p.Right != nil {
				queueP = append(queueP, p.Right)
				queueQ = append(queueQ, q.Right)
			}
		}
	}
	return true
}

package main

/*
  二叉搜索树中的两个节点被错误地交换。
  请在不改变其结构的情况下，恢复这棵树。
*/

// Morris Traversal
func recoverTree(root *TreeNode) {
	var (
		cur    = root
		pre    *TreeNode
		first  *TreeNode
		second *TreeNode
	)
	handle := func() {
		if pre != nil && pre.Val > cur.Val {
			if first == nil {
				first = pre
			}
			second = cur
		}
		pre = cur
		cur = cur.Right
	}
	for cur != nil {
		if cur.Left != nil {
			curPre := cur.Left
			for curPre.Right != nil && curPre.Right != cur {
				curPre = curPre.Right
			}
			if curPre.Right != cur {
				curPre.Right, cur = cur, cur.Left
				continue
			}
			curPre.Right = nil
			handle()
		} else {
			handle()
		}
	}
	first.Val, second.Val = second.Val, first.Val
}

// Stack
func recoverTree2(root *TreeNode) {
	var (
		node       = root
		stack      []*TreeNode
		preNode    *TreeNode
		curNode    *TreeNode
		firstNode  *TreeNode
		secondNode *TreeNode
	)
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) != 0 {
			preNode, curNode = curNode, stack[len(stack)-1]
			if preNode != nil && preNode.Val > curNode.Val {
				if firstNode == nil {
					firstNode = preNode
				}
				secondNode = curNode
			}
			node = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}

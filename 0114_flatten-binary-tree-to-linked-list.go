package leetcode

/*
  给定一个二叉树，原地将它展开为链表。
  例如，给定二叉树
        1
       / \
      2   5
     / \   \
    3   4   6
  将其展开为：
    1
     \
      2
       \
        3
         \
          4
           \
            5
             \
              6
*/

// Recursive
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		cur := root.Left
		for cur.Right != nil {
			cur = cur.Right
		}
		right := root.Right
		root.Right = root.Left
		cur.Right = right
		root.Left = nil
	}
	next := root.Right
	for next != nil && next.Left == nil {
		next = next.Right
	}
	flatten(next)
}

// Morris Traversal
func flatten2(root *TreeNode) {
	for cur := root; cur != nil; {
		if cur.Left != nil {
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right != cur {
				pre.Right, cur = cur, cur.Left
				continue
			}
			pre.Right, cur.Right, cur.Left = cur.Right, cur.Left, nil
		} else {
			cur = cur.Right
		}
	}
}

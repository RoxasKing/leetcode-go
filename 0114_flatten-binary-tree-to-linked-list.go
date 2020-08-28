package main

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

// Preorder Traversal (VLR)
func flatten(root *TreeNode) {
	for cur := root; cur != nil; cur = cur.Right {
		if cur.Left != nil {
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
	}
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

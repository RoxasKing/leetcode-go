package leetcode

/*
  给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
  本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
  示例:
    给定的有序链表： [-10, -3, 0, 5, 9],
    一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
          0
         / \
       -3   9
       /   /
     -10  5
*/

func sortedListToBST(head *ListNode) *TreeNode {
	var buildBST func(*ListNode) *TreeNode
	buildBST = func(head *ListNode) *TreeNode {
		if head == nil {
			return nil
		} else if head.Next == nil {
			return &TreeNode{Val: head.Val}
		}
		slowPre := &ListNode{Val: -1, Next: head}
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slowPre = slowPre.Next
			slow, fast = slow.Next, fast.Next.Next
		}
		slowPre.Next = nil
		return &TreeNode{
			Val:   slow.Val,
			Left:  buildBST(head),
			Right: buildBST(slow.Next),
		}
	}
	return buildBST(head)
}

// Convert linked list to array
func sortedListToBST2(head *ListNode) *TreeNode {
	var array []int
	for node := head; node != nil; node = node.Next {
		array = append(array, node.Val)
	}
	var buildBST func(int, int) *TreeNode
	buildBST = func(left, right int) *TreeNode {
		if left == right {
			return nil
		}
		mid := left + (right-left)>>1
		return &TreeNode{
			Val:   array[mid],
			Left:  buildBST(left, mid),
			Right: buildBST(mid+1, right),
		}
	}
	return buildBST(0, len(array))
}

// Inorder traversal simulation
func sortedListToBST3(head *ListNode) *TreeNode {
	var size int
	for node := head; node != nil; node = node.Next {
		size++
	}
	var buildBST func(int, int) *TreeNode
	buildBST = func(l, r int) *TreeNode {
		if l == r {
			return nil
		}
		mid := l + (r-l)>>1
		left := buildBST(l, mid)
		node := &TreeNode{Val: head.Val}
		node.Left = left
		head = head.Next
		node.Right = buildBST(mid+1, r)
		return node
	}
	return buildBST(0, size)
}

package codinginterviews

/*
  输入两个链表，找出它们的第一个公共节点。

  注意：
    如果两个链表没有交点，返回 null.
    在返回结果后，两个链表仍须保持原有的结构。
    可假定整个链表结构中没有循环。
    程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lenA, lenB int
	for node := headA; node != nil; node = node.Next {
		lenA++
	}
	for node := headB; node != nil; node = node.Next {
		lenB++
	}
	ptrA := headA
	ptrB := headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			ptrA = ptrA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			ptrB = ptrB.Next
		}
	}
	for ptrA != nil {
		if ptrA == ptrB {
			return ptrA
		}
		ptrA, ptrB = ptrA.Next, ptrB.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ptrA, ptrB := headA, headB
	for ptrA != ptrB {
		if ptrA != nil {
			ptrA = ptrA.Next
		} else {
			ptrA = headB
		}
		if ptrB != nil {
			ptrB = ptrB.Next
		} else {
			ptrB = headA
		}
	}
	return ptrA
}

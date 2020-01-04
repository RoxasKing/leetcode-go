package My_LeetCode_In_Go

/*
  合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	base := 1
	newlists := make([]*ListNode, len(lists))
	copy(newlists, lists)
	for base <= len(lists) {
		for i := 0; i < len(newlists); i += 2 * base {
			if i+base < len(lists) {
				newlists[i] = mergeTwoLists(newlists[i], newlists[i+base])
			}
		}
		base *= 2
	}
	return newlists[0]
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge_k_lists(&lists, 0, len(lists)-1)
}

func merge_k_lists(lists *[]*ListNode, l, r int) *ListNode {
	if l == r {
		return (*lists)[l]
	}
	if l > r {
		return nil
	}
	mid := l + ((r - l) >> 1)
	return mergeTwoLists(
		merge_k_lists(lists, l, mid),
		merge_k_lists(lists, mid+1, r),
	)
}

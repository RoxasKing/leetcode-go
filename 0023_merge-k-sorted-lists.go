package leetcode

/*
  合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	newLists := make([]*ListNode, len(lists))
	copy(newLists, lists)
	for base := 1; base <= len(lists); base *= 2 {
		for i := 0; i < len(newLists); i += 2 * base {
			if i+base < len(lists) {
				newLists[i] = mergeTwoLists(newLists[i], newLists[i+base])
			}
		}
	}
	return newLists[0]
}

// Recursive
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return recurMergeKLists(&lists, 0, len(lists)-1)
}

func recurMergeKLists(lists *[]*ListNode, l, r int) *ListNode {
	if l == r {
		return (*lists)[l]
	}
	m := l + (r-l)>>1
	return mergeTwoLists(recurMergeKLists(lists, l, m), recurMergeKLists(lists, m+1, r))
}

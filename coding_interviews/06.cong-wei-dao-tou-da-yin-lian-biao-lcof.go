package codinginterviews

/*
  输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

  限制：
  0 <= 链表长度 <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reversePrint(head *ListNode) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	for i := 0; i < len(out)>>1; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return out
}

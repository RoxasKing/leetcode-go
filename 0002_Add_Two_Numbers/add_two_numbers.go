package _002_Add_Two_Numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 头节点
	resPre := &ListNode{}
	// 轮询节点
	cur := resPre
	// 进位数
	carry := 0

	// 只要链表其一有数，或者进位数不为 0 ，循环继续
	for l1 != nil || l2 != nil || carry > 0 {
		// 初始化下一节点起始数
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}

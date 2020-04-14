package leetcode

/*
  给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
  你可以假设除了数字 0 之外，这两个数字都不会以零开头。
  进阶：
  如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
*/

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1 []*ListNode
	for l1 != nil {
		stack1 = append(stack1, l1)
		l1 = l1.Next
	}
	var stack2 []*ListNode
	for l2 != nil {
		stack2 = append(stack2, l2)
		l2 = l2.Next
	}
	var head *ListNode
	var rest int
	for len(stack1) != 0 || len(stack2) != 0 {
		if len(stack1) != 0 {
			rest += stack1[len(stack1)-1].Val
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			rest += stack2[len(stack2)-1].Val
			stack2 = stack2[:len(stack2)-1]
		}
		node := &ListNode{Val: rest % 10, Next: head}
		head = node
		rest /= 10
	}
	if rest != 0 {
		node := &ListNode{Val: rest, Next: head}
		head = node
	}
	return head
}

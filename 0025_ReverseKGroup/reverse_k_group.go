package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var out *ListNode
	// 要翻转节点的前一个节点
	pre := &ListNode{}
	pre.Next = head
	out = pre

	for {
		i := 0
		end := head
		// 测试剩余节点数量是否满足翻转条件
		for ; i < k-1; i++ {
			if end != nil {
				end = end.Next
			} else {
				break
			}
		}
		if i < k-1 || end == nil {
			// 不满足 退出
			return out.Next
		} else {
			// 满足 进行翻转

			// 需要翻转的节点的后继节点
			next := end.Next

			tmp := end
			// 针对需要翻转的节点处理流程
			for i := 0; i < k-2; i++ {
				node := head
				// 取最后一个节点的上一个节点
				for node.Next != tmp {
					node = node.Next
				}
				// 将节点插入到末位节点之后
				tmp.Next = node
				tmp = tmp.Next
			}
			// 插入头节点
			tmp.Next = head
			// 头节点指向后继节点
			head.Next = next
			// 前节点指向末位节点，完成翻转操作
			pre.Next = end

			// 为下一次翻转做数据准备
			pre = head
			head = pre.Next
		}
	}
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 将链表存入栈中
	var stack []*ListNode
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}
	// 获取链表总长度
	n := len(stack)
	// 建立链表头指针
	head = &ListNode{Val: -1, Next: nil}
	// 初始化尾节点
	tail := head
	// 计算剩余不够一组的数量
	j := n % k
	// 将这些节点正常顺序插入到头节点之后
	for i := 0; i < j; i++ {
		// 获取栈顶节点
		node := stack[n-1]
		// 新节点连接头指针的后继节点
		node.Next = head.Next
		// 头指针连接新节点
		head.Next = node
		// 尾指针指向新节点
		tail = node
		n--
	}
	// 翻转节点流程
	for n > 0 {
		// 获取栈顶节点
		node := stack[n-1]
		// 新节点连接头指针的后继节点
		node.Next = head.Next
		// 头指针连接新节点
		head.Next = node
		// 尾指针指向新节点
		tail = node
		n--
		// 翻转当前组节点
		for i := 0; i < k-1; i++ {
			// 获取栈顶节点
			node := stack[n-1]
			// 新节点连接尾指针的后继节点
			node.Next = tail.Next
			// 尾指针连接新节点
			tail.Next = node
			// 尾指针指向新节点
			tail = node
			n--
		}
	}
	return head.Next
}

func reverseKGroup3(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 新建链表
	pre := &ListNode{Val: -1, Next: nil}
	// 初始化尾节点
	tail := pre
	for {
		// 定义一个链表数组 , 当作栈用
		var stack []*ListNode
		i := 0
		// 将前 k 个节点放入数组中
		for i < k && head != nil {
			stack = append(stack, head)
			head = head.Next
			i++
		}
		if i == k {
			// 如果数组大小满足 k 个 , 进行翻转
			for j := 0; j < i; j++ {
				// 获取栈顶元素
				node := stack[len(stack)-1]
				// 将新元素插入到尾指针之后
				node.Next = tail.Next
				// 尾指针连接新节点
				tail.Next = node
				// 尾指针移动到新节点上
				tail = node
				// 数组大小减 1 位
				stack = stack[:len(stack)-1]
			}
		} else {
			// 不满足 k 个的 ， 按原样输出
			for _, node := range stack {
				tail.Next = node
				tail = tail.Next
			}
			return pre.Next
		}
	}
}

func main() {
	head1 := &ListNode{Val: 1}
	nex1 := &ListNode{Val: 2}
	head1.Next = nex1
	nex2 := &ListNode{Val: 3}
	nex1.Next = nex2
	nex3 := &ListNode{Val: 4}
	nex2.Next = nex3
	nex4 := &ListNode{Val: 5}
	nex3.Next = nex4

	new := reverseKGroup3(head1, 2)

	for node := new; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}

}

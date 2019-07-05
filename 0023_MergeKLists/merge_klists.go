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

func mergeKLists2(lists []*ListNode) *ListNode {
	var out *ListNode
	var tmp *ListNode
	var array []int

	// 先放入同一个数组中
	for i := 0; i < len(lists); i++ {
		for node := lists[i]; node != nil; node = node.Next {
			array = append(array, node.Val)
		}
	}
	// 并排序
	QuickSort(array)

	// 生成排序后的链表
	for j := 0; j < len(array); j++ {
		// 初始链表为空时
		if out == nil {
			out = &ListNode{}
			tmp = out
			tmp.Val = array[j]
		} else if tmp != nil {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
			tmp.Val = array[j]
		}
	}
	return out
}

func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}

func main() {
	var lists []*ListNode
	//var head1 *ListNode
	//head1 := &ListNode{Val: 1}
	//nex11 := &ListNode{Val: 4}
	//head1.Next = nex11
	//head2 := &ListNode{Val: 2}
	//head3 := &ListNode{Val: 3}
	//lists = append(lists, head1)
	//lists = append(lists, head1, head2, head3)
	//for i := 0; i < len(lists); i++ {
	//	fmt.Println(lists[i].Val)
	//}
	for elem := mergeKLists2(lists); elem != nil; elem = elem.Next {
		fmt.Println(elem.Val)
	}
}

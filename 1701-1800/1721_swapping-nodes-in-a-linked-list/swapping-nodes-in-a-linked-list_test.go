package main

import (
	"reflect"
	"testing"
)

func Test_swapNodes(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5}, 2}, []int{1, 4, 3, 2, 5}},
		{"2", args{[]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}, 5}, []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5}},
		{"3", args{[]int{1}, 1}, []int{1}},
		{"4", args{[]int{1, 2}, 1}, []int{2, 1}},
		{"5", args{[]int{1, 2, 3}, 2}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(arr2List(tt.args.arr), tt.args.k); !reflect.DeepEqual(list2Arr(got), tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func arr2List(arr []int) *ListNode {
	var head, tail *ListNode
	for _, num := range arr {
		if head == nil {
			head = &ListNode{Val: num}
			tail = head
		} else {
			tail.Next = &ListNode{Val: num}
			tail = tail.Next
		}
	}
	return head
}

func list2Arr(head *ListNode) []int {
	var out []int
	for ; head != nil; head = head.Next {
		out = append(out, head.Val)
	}
	return out
}

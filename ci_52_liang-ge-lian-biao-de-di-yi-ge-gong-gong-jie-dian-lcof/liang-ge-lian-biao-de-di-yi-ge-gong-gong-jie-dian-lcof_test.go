package main

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	a1 := &ListNode{Val: 4}
	a2 := &ListNode{Val: 1}
	b1 := &ListNode{Val: 5}
	b2 := &ListNode{Val: 0}
	b3 := &ListNode{Val: 1}
	c1 := &ListNode{Val: 8}
	c2 := &ListNode{Val: 4}
	c3 := &ListNode{Val: 5}
	a1.Next = a2
	b1.Next = b2
	b2.Next = b3
	c1.Next = c2
	c2.Next = c3
	a2.Next = c1
	b3.Next = c1
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{a1, b1}, c1},
		{"2", args{nil, nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

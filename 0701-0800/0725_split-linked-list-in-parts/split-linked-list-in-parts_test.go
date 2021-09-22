package main

import (
	"reflect"
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3

	n4 := &ListNode{Val: 1}
	n5 := &ListNode{Val: 2}
	n6 := &ListNode{Val: 3}
	n7 := &ListNode{Val: 4}
	n8 := &ListNode{Val: 5}
	n9 := &ListNode{Val: 6}
	n10 := &ListNode{Val: 7}
	n11 := &ListNode{Val: 8}
	n12 := &ListNode{Val: 9}
	n13 := &ListNode{Val: 10}
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	n7.Next = n8
	n8.Next = n9
	n9.Next = n10
	n10.Next = n11
	n11.Next = n12
	n12.Next = n13
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		{
			"1",
			args{n1, 5},
			[]*ListNode{{Val: 1}, {Val: 2}, {Val: 3}, nil, nil},
		},
		{
			"2",
			args{n4, 3},
			[]*ListNode{
				{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
				{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}},
				{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitListToParts(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

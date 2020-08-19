package main

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	n1 := &ListNode{Val: 4}
	n2 := &ListNode{Val: 5}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 9}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{
				&ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  1,
							Next: &ListNode{Val: 9},
						},
					},
				},
			},
			&ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 9},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
			if !reflect.DeepEqual(tt.args.node, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", tt.args.node, tt.want)
			}
		})
	}
}

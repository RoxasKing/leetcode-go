package main

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{
				[]*ListNode{
					nil,
					{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 11}}},
					nil,
					{Val: 6, Next: &ListNode{Val: 10}},
				},
			},
			&ListNode{
				Val: -1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  10,
							Next: &ListNode{Val: 11},
						},
					},
				},
			},
		},
		{
			"1",
			args{
				[]*ListNode{
					{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
					{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
					{Val: 2, Next: &ListNode{Val: 6}},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val:  5,
										Next: &ListNode{Val: 6},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(tt.args.lists)
			for node := tt.want; node != nil; node, got = node.Next, got.Next {
				if node.Val != got.Val {
					t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

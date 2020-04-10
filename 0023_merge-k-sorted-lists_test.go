package leetcode

import (
	"reflect"
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
			"",
			args{
				[]*ListNode{
					nil,
					{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 11, Next: nil}}},
					nil,
					{Val: 6, Next: &ListNode{Val: 10, Next: nil}},
				},
			},
			&ListNode{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 10, Next: &ListNode{Val: 11, Next: nil}}}}},
		},
		{
			"",
			args{
				[]*ListNode{
					{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
					{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
					{Val: 2, Next: &ListNode{Val: 6, Next: nil}},
				},
			},
			&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}}},
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

func Test_mergeKLists2(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"",
			args{
				[]*ListNode{
					nil,
					{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 11, Next: nil}}},
					nil,
					{Val: 6, Next: &ListNode{Val: 10, Next: nil}},
				},
			},
			&ListNode{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 10, Next: &ListNode{Val: 11, Next: nil}}}}},
		},
		{
			"",
			args{
				[]*ListNode{
					{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
					{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
					{Val: 2, Next: &ListNode{Val: 6, Next: nil}},
				},
			},
			&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists2(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists2() = %v, want %v", got, tt.want)
			}
		})
	}
}

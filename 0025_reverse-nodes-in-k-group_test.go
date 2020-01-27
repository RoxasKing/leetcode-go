package My_LeetCode_In_Go

import (
	"testing"

	. "My_LeetCode_In_Go/util/linkedlist"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"",
			args{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}, 2},
			&ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}},
		},
		{
			"",
			args{
				&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
				2,
			},
			&ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}}},
		},
		{
			"",
			args{
				&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
				3,
			},
			&ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			for node := tt.want; node != nil; node, got = node.Next, got.Next {
				if node.Val != got.Val {
					t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

package My_LeetCode_In_Go

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test 1",
			args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
			&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(tt.args.head)
			for node := tt.want; node != nil; node, got = node.Next, got.Next {
				if node.Val != got.Val {
					t.Errorf("swapPairs() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

package My_LeetCode_In_Go

import (
	"testing"
)

func Test_sortList(t *testing.T) {
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
			args{
				&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}},
			},
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortList(tt.args.head)
			for l1, l2 := tt.want, got; l1 != nil; {
				if l1.Val != l2.Val {
					t.Errorf("sortList() = %v, want %v", got, tt.want)
				}
				l1, l2 = l1.Next, l2.Next
			}
		})
	}
}

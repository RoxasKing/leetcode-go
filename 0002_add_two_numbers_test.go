package My_LeetCode_In_Go

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"",
			args{
				&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
				&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			for got != nil {
				if got.Val != tt.want.Val {
					t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
				}
				got = got.Next
				tt.want = tt.want.Next
			}
		})
	}
}

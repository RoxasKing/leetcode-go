package leetcode

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{
				&ListNode{Val: 1},
				1,
			},
			nil,
		},
		{
			"2",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{Val: 5},
							},
						},
					},
				},
				2,
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 5},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			for ; got != nil; got, tt.want = got.Next, tt.want.Next {
				if got.Val != tt.want.Val {
					t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

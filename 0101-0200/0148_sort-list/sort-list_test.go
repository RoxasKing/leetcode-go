package main

import "testing"

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
			"1",
			args{
				&ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: &ListNode{Val: 3},
						},
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 4},
					},
				},
			},
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

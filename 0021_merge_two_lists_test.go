package My_LeetCode_In_Go

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
				&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
				&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			},
			&ListNode{1,
				&ListNode{1,
					&ListNode{2,
						&ListNode{3,
							&ListNode{4,
								&ListNode{4, nil},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.l1, tt.args.l2)
			for node := tt.want; node != nil; node, got = node.Next, got.Next {
				if node.Val != got.Val {
					t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

package My_LeetCode_In_Go

import (
	"testing"
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
			args{&ListNode{1, &ListNode{2, nil}}, 2},
			&ListNode{2, &ListNode{1, nil}},
		},
		{
			"",
			args{
				&ListNode{1,
					&ListNode{2,
						&ListNode{3,
							&ListNode{4,
								&ListNode{5, nil},
							},
						},
					},
				},
				2,
			},
			&ListNode{2,
				&ListNode{1,
					&ListNode{4,
						&ListNode{3,
							&ListNode{5, nil},
						},
					},
				},
			},
		},
		{
			"",
			args{
				&ListNode{1,
					&ListNode{2,
						&ListNode{3,
							&ListNode{4,
								&ListNode{5, nil},
							},
						},
					},
				},
				3,
			},
			&ListNode{3,
				&ListNode{2,
					&ListNode{1,
						&ListNode{4,
							&ListNode{5, nil},
						},
					},
				},
			},
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

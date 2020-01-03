package My_LeetCode_In_Go

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
		{"test 2", args{&ListNode{1, nil}, 1}, nil},
		{
			"test 1",
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
			&ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{5, nil},
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

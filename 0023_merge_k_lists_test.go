package My_LeetCode_In_Go

import (
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
					&ListNode{-1, &ListNode{5, &ListNode{11, nil}}},
					nil,
					&ListNode{6, &ListNode{10, nil}},
				},
			},
			&ListNode{-1,
				&ListNode{5,
					&ListNode{6,
						&ListNode{10,
							&ListNode{11, nil},
						},
					},
				},
			},
		},
		{
			"",
			args{
				[]*ListNode{
					&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
					&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
					&ListNode{2, &ListNode{6, nil}},
				},
			},
			&ListNode{1,
				&ListNode{1,
					&ListNode{2,
						&ListNode{3,
							&ListNode{4,
								&ListNode{4,
									&ListNode{5,
										&ListNode{6, nil},
									},
								},
							},
						},
					},
				},
			},
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

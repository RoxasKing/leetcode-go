package leetcode

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"", args{&ListNode{Val: 1, Next: &ListNode{Val: 1}}}, nil},
		{
			"",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 3,
								},
							},
						},
					},
				},
			},
			&ListNode{
				Val:  2,
				Next: &ListNode{Val: 3, Next: nil},
			},
		},
		{
			"",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val:  4,
										Next: &ListNode{Val: 5, Next: nil},
									},
								},
							},
						},
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{Val: 5, Next: nil},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates0082(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

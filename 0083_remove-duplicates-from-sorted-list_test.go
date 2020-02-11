package leetcode

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates0083(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  3,
								Next: &ListNode{Val: 3},
							},
						},
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{Val: 3},
				},
			},
		},
		{
			"",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: &ListNode{Val: 2},
					},
				},
			},
			&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates0083(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates0083() = %v, want %v", got, tt.want)
			}
		})
	}
}

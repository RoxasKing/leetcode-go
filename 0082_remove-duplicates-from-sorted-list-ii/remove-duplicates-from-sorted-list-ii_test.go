package main

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
		{
			"1",
			args{
				&ListNode{
					Val:  1,
					Next: &ListNode{Val: 1},
				},
			},
			nil,
		},
		{
			"2",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val:  2,
								Next: &ListNode{Val: 3},
							},
						},
					},
				},
			},
			&ListNode{
				Val:  2,
				Next: &ListNode{Val: 3},
			},
		},
		{
			"3",
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
										Next: &ListNode{Val: 5},
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
					Next: &ListNode{Val: 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

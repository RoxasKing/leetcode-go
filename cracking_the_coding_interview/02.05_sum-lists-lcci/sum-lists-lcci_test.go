package main

import (
	"reflect"
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
			"1",
			args{
				&ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  1,
						Next: &ListNode{Val: 6},
					},
				},
				&ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  9,
						Next: &ListNode{Val: 2},
					},
				},
			},
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 9},
				},
			},
		},
		{
			"2",
			args{
				&ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  1,
						Next: &ListNode{Val: 7},
					},
				},
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  9,
						Next: &ListNode{Val: 5},
					},
				},
			},
			&ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 1},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
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
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{Val: 5, Next: nil},
							},
						},
					},
				},
				1,
			},
			&ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{Val: 4, Next: nil},
						},
					},
				},
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
								Val:  4,
								Next: &ListNode{Val: 5, Next: nil},
							},
						},
					},
				},
				2,
			},
			&ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{Val: 3, Next: nil},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

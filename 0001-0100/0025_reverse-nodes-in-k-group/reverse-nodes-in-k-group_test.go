package main

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
	}{
		{
			"1",
			args{
				&ListNode{
					Val:  1,
					Next: &ListNode{Val: 2},
				},
				2,
			},
		},
		{
			"2",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{Val: 5},
							},
						},
					},
				},
				2,
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
								Val:  4,
								Next: &ListNode{Val: 5},
							},
						},
					},
				},
				3,
			},
		},
		{
			"4",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{Val: 5},
							},
						},
					},
				},
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr, k := []*ListNode{}, tt.args.k
			for n := tt.args.head; n != nil; n = n.Next {
				arr = append(arr, n)
			}
			for i := 0; i+k-1 < len(arr); i += k {
				reverse(arr[i : i+k])
			}
			got := reverseKGroup(tt.args.head, tt.args.k)
			for n := got; n != nil; n = n.Next {
				if n != arr[0] {
					t.Errorf("reverseKGroup() test failed!")
					return
				}
				arr = arr[1:]
			}
		})
	}
}

func reverse(arr []*ListNode) {
	n := len(arr)
	for i := 0; i < n>>1; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

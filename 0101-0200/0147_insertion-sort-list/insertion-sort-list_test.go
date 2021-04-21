package main

import (
	"sort"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				&ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
		},
		{
			"2",
			args{
				&ListNode{
					Val: -1,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 0,
								},
							},
						},
					},
				},
			},
		},
		{"3", args{nil}},
		{"4", args{&ListNode{Val: -1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var array []*ListNode
			for n := tt.args.head; n != nil; n = n.Next {
				array = append(array, n)
			}
			sort.Slice(array, func(i, j int) bool {
				return array[i].Val < array[j].Val
			})
			got := insertionSortList(tt.args.head)
			for n := got; n != nil; n = n.Next {
				if len(array) == 0 || n.Val != array[0].Val {
					t.Error("insertionSortList() test failed")
					return
				}
				array = array[1:]
			}
			if len(array) != 0 {
				t.Error("insertionSortList() test failed")
			}
		})
	}
}

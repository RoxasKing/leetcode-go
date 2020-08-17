package main

import (
	"reflect"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"1",
			args{
				&ListNode{
					Val: -10,
					Next: &ListNode{
						Val: -3,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val:  5,
								Next: &ListNode{Val: 9},
							},
						},
					},
				},
			},
			&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedListToBST2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"1",
			args{
				&ListNode{
					Val: -10,
					Next: &ListNode{
						Val: -3,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val:  5,
								Next: &ListNode{Val: 9},
							},
						},
					},
				},
			},
			&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedListToBST3(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"1",
			args{
				&ListNode{
					Val: -10,
					Next: &ListNode{
						Val: -3,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val:  5,
								Next: &ListNode{Val: 9},
							},
						},
					},
				},
			},
			&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST3(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST3() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{&ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 9},
				},
			}},
			&ListNode{
				Val:  1,
				Next: &ListNode{Val: 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
		})
	}
}

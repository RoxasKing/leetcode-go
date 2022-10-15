package main

import "testing"

func Test_numComponents(t *testing.T) {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 2}
	l4 := &ListNode{Val: 3}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	type args struct {
		head *ListNode
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{l1, []int{0, 1, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numComponents(tt.args.head, tt.args.nums); got != tt.want {
				t.Errorf("numComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

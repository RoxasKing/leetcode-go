package leetcode

import (
	"testing"
)

func Test_hasCycle(t *testing.T) {
	node1 := &ListNode{Val: 1, Next: nil}
	node1.Next = &ListNode{Val: 2, Next: nil}
	node1.Next.Next = &ListNode{Val: 3, Next: node1}
	node2 := &ListNode{Val: 1, Next: nil}
	node2.Next = &ListNode{Val: 2, Next: nil}
	node2.Next.Next = &ListNode{Val: 3, Next: nil}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{node1}, true},
		{"2", args{node2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

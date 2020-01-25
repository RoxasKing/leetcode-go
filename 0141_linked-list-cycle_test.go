package My_LeetCode_In_Go

import (
	"testing"
)

func Test_hasCycle(t *testing.T) {
	node1 := &ListNode{1, nil}
	node1.Next = &ListNode{2, nil}
	node1.Next.Next = &ListNode{3, node1}
	node2 := &ListNode{1, nil}
	node2.Next = &ListNode{2, nil}
	node2.Next.Next = &ListNode{3, nil}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{node1}, true},
		{"", args{node2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

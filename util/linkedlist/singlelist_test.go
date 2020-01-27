package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"", args{&ListNode{0, &ListNode{1, &ListNode{2, nil}}}},
			&ListNode{2, &ListNode{1, &ListNode{0, nil}}}},
		{"", args{&ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}},
			&ListNode{3, &ListNode{2, &ListNode{1, &ListNode{0, nil}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasCycle(t *testing.T) {
	head1 := &ListNode{0, nil}
	head1.Next = &ListNode{1, nil}
	head1.Next.Next = &ListNode{2, nil}
	head1.Next.Next.Next = head1
	head2 := &ListNode{0, nil}
	head2.Next = &ListNode{1, nil}
	head2.Next.Next = &ListNode{2, nil}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{head1}, true},
		{"", args{head2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle(tt.args.head); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

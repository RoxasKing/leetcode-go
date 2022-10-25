package main

import (
	"reflect"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
	a1 := &ListNode{1, nil}
	a2 := &ListNode{3, nil}
	a3 := &ListNode{4, nil}
	a4 := &ListNode{7, nil}
	a5 := &ListNode{1, nil}
	a6 := &ListNode{2, nil}
	a7 := &ListNode{6, nil}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	a5.Next = a6
	a6.Next = a7
	b1 := &ListNode{1, nil}
	b2 := &ListNode{3, nil}
	b3 := &ListNode{4, nil}
	b5 := &ListNode{1, nil}
	b6 := &ListNode{2, nil}
	b7 := &ListNode{6, nil}
	b1.Next = b2
	b2.Next = b3
	b3.Next = b5
	b5.Next = b6
	b6.Next = b7

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{a1}, b1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

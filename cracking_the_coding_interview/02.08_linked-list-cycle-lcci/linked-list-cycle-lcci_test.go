package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	a1 := &ListNode{Val: 3}
	a2 := &ListNode{Val: 2}
	a3 := &ListNode{Val: 0}
	a4 := &ListNode{Val: -4}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a2

	b1 := &ListNode{Val: 1}
	b2 := &ListNode{Val: 2}
	b1.Next = b2
	b2.Next = b1

	c1 := &ListNode{Val: 1}

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{a1},
			a2,
		},
		{
			"2",
			args{b1},
			b1,
		},
		{
			"3",
			args{c1},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

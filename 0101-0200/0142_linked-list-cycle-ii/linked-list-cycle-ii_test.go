package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	n0 := &ListNode{Val: 3}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 0}
	n3 := &ListNode{Val: -4}
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n1
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{n0}, n1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

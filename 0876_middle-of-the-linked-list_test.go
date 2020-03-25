package leetcode

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	head := &ListNode{Val: 1}
	next1 := &ListNode{Val: 2}
	next2 := &ListNode{Val: 3}
	next3 := &ListNode{Val: 4}
	next4 := &ListNode{Val: 5}
	next5 := &ListNode{Val: 6}
	head.Next = next1
	next1.Next = next2
	next2.Next = next3
	next3.Next = next4
	next4.Next = next5
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"",
			args{head},
			next3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

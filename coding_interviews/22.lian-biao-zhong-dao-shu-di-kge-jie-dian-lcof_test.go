package codinginterviews

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 4}
	e := &ListNode{Val: 5}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{a, 2}, d},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

package codinginterviews

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 4}
	e := &ListNode{Val: 5}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	f := &ListNode{Val: 5}
	g := &ListNode{Val: 4}
	h := &ListNode{Val: 3}
	i := &ListNode{Val: 2}
	j := &ListNode{Val: 1}
	f.Next = g
	g.Next = h
	h.Next = i
	i.Next = j
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{a}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

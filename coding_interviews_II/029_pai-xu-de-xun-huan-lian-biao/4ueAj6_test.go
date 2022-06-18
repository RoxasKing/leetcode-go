package main

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	a1 := &Node{Val: 3}
	a2 := &Node{Val: 4}
	a3 := &Node{Val: 1}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a1

	b1 := &Node{Val: 3}
	b2 := &Node{Val: 4}
	b3 := &Node{Val: 1}
	b4 := &Node{Val: 2}
	b1.Next = b2
	b2.Next = b3
	b3.Next = b4
	b4.Next = b1

	c1 := &Node{Val: 1}
	c1.Next = c1

	d1 := &Node{Val: 1}
	d1.Next = d1

	e1 := &Node{Val: 1}
	e2 := &Node{Val: 0}
	e1.Next = e2
	e2.Next = e1

	f1 := &Node{Val: 1}
	f2 := &Node{Val: 2}
	f1.Next = f2
	f2.Next = f1

	g1 := &Node{Val: 1}
	g2 := &Node{Val: 2}
	g3 := &Node{Val: 3}
	g1.Next = g2
	g2.Next = g3
	g3.Next = g1

	type args struct {
		aNode *Node
		x     int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"1", args{a1, 2}, b1},
		{"2", args{nil, 1}, c1},
		{"3", args{d1, 0}, e1},
		{"4", args{f1, 3}, g1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.aNode, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

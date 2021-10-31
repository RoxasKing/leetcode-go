package main

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	p1 := &Node{Val: 1}
	p2 := &Node{Val: 2}
	p3 := &Node{Val: 3}
	p4 := &Node{Val: 4}
	p5 := &Node{Val: 5}
	p6 := &Node{Val: 6}
	p7 := &Node{Val: 7}
	p8 := &Node{Val: 8}
	p11 := &Node{Val: 11}
	p12 := &Node{Val: 12}
	p1.Next = p2
	p2.Prev = p1
	p2.Next = p3
	p3.Prev = p2
	p3.Next = p4
	p4.Prev = p3
	p4.Next = p5
	p5.Prev = p4
	p5.Next = p6
	p6.Prev = p5
	p3.Child = p7
	p7.Next = p8
	p8.Prev = p7
	p8.Child = p11
	p11.Next = p12
	p12.Prev = p11

	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n11 := &Node{Val: 11}
	n12 := &Node{Val: 12}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n1.Next = n2
	n2.Prev = n1
	n2.Next = n3
	n3.Prev = n2
	n3.Next = n7
	n7.Prev = n3
	n7.Next = n8
	n8.Prev = n7
	n8.Next = n11
	n11.Prev = n8
	n11.Next = n12
	n12.Prev = n11
	n12.Next = n4
	n4.Prev = n12
	n4.Next = n5
	n5.Prev = n4
	n5.Next = n6
	n6.Prev = n5

	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"1", args{p1}, n1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatten(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

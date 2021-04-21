package main

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	n0 := &Node{Val: 1}
	n1 := &Node{Val: 2}
	n2 := &Node{Val: 3}
	n0.Left = n1
	n0.Right = n2
	n3 := &Node{Val: 4}
	n4 := &Node{Val: 5}
	n1.Left = n3
	n1.Right = n4
	n5 := &Node{Val: 7}
	n2.Right = n5

	n0c := &Node{Val: 1}
	n1c := &Node{Val: 2}
	n2c := &Node{Val: 3}
	n0c.Left = n1c
	n0c.Right = n2c
	n1c.Next = n2c
	n3c := &Node{Val: 4}
	n4c := &Node{Val: 5}
	n1c.Left = n3c
	n1c.Right = n4c
	n3c.Next = n4c
	n5c := &Node{Val: 7}
	n2c.Right = n5c
	n4c.Next = n5c
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"1",
			args{
				n0,
			},
			n0c,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

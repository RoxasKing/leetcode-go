package main

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	ex1 := &Node{Val: 7}
	ex11 := &Node{Val: 13}
	ex12 := &Node{Val: 11}
	ex13 := &Node{Val: 10}
	ex14 := &Node{Val: 1}
	ex1.Next = ex11
	ex11.Next = ex12
	ex11.Random = ex1
	ex12.Next = ex13
	ex12.Random = ex14
	ex13.Next = ex14
	ex13.Random = ex12
	ex14.Random = ex1

	ex2 := &Node{Val: 1}
	ex21 := &Node{Val: 2}
	ex2.Next = ex21
	ex21.Random = ex21

	ex3 := &Node{Val: 3}
	ex31 := &Node{Val: 3}
	ex32 := &Node{Val: 3}
	ex3.Next = ex31
	ex31.Next = ex32
	ex31.Random = ex3

	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"1", args{ex1}, ex1},
		{"2", args{ex2}, ex2},
		{"3", args{ex3}, ex3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}

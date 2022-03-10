package main

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 3}
	n3 := &Node{Val: 2}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n1.Children = append(n1.Children, n2, n3, n4)
	n2.Children = append(n2.Children, n5, n6)

	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{n1}, []int{1, 3, 5, 6, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

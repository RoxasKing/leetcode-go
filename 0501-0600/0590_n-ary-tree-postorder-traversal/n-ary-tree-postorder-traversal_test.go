package main

import (
	"reflect"
	"testing"
)

func Test_postorder(t *testing.T) {
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
		{"1", args{n1}, []int{5, 6, 3, 2, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

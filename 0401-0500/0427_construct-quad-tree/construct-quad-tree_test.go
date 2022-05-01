package main

import (
	"reflect"
	"testing"
)

func Test_construct(t *testing.T) {
	a0 := &Node{Val: true, IsLeaf: false}
	a1 := &Node{Val: false, IsLeaf: true}
	a2 := &Node{Val: true, IsLeaf: true}
	a3 := &Node{Val: true, IsLeaf: true}
	a4 := &Node{Val: false, IsLeaf: true}
	a0.TopLeft = a1
	a0.TopRight = a2
	a0.BottomLeft = a3
	a0.BottomRight = a4

	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"1", args{[][]int{{0, 1}, {1, 0}}}, a0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct() = %v, want %v", got, tt.want)
			}
		})
	}
}

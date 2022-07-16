package main

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		quadTree1 *Node
		quadTree2 *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"1",
			args{
				&Node{
					Val:         true,
					IsLeaf:      false,
					TopLeft:     &Node{Val: true, IsLeaf: true},
					TopRight:    &Node{Val: true, IsLeaf: true},
					BottomLeft:  &Node{Val: false, IsLeaf: true},
					BottomRight: &Node{Val: false, IsLeaf: true},
				},
				&Node{
					Val:     true,
					IsLeaf:  false,
					TopLeft: &Node{Val: true, IsLeaf: true},
					TopRight: &Node{
						Val:         true,
						IsLeaf:      false,
						TopLeft:     &Node{Val: false, IsLeaf: true},
						TopRight:    &Node{Val: false, IsLeaf: true},
						BottomLeft:  &Node{Val: true, IsLeaf: true},
						BottomRight: &Node{Val: true, IsLeaf: true},
					},
					BottomLeft:  &Node{Val: true, IsLeaf: true},
					BottomRight: &Node{Val: false, IsLeaf: true},
				},
			},
			&Node{
				Val:         false,
				IsLeaf:      false,
				TopLeft:     &Node{Val: true, IsLeaf: true},
				TopRight:    &Node{Val: true, IsLeaf: true},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: false, IsLeaf: true},
			},
		},
		{
			"2",
			args{&Node{Val: false, IsLeaf: true}, &Node{Val: false, IsLeaf: true}},
			&Node{Val: false, IsLeaf: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.quadTree1, tt.args.quadTree2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

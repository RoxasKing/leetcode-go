package main

import "testing"

func Test_maxDepth(t *testing.T) {
	a1 := &Node{Val: 1}
	a2 := &Node{Val: 2}
	a3 := &Node{Val: 3}
	a4 := &Node{Val: 4}
	a5 := &Node{Val: 5}
	a6 := &Node{Val: 6}
	a1.Children = []*Node{a2, a3, a4}
	a3.Children = []*Node{a5, a6}

	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{a1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

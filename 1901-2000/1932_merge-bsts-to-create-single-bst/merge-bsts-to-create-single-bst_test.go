package main

import (
	"reflect"
	"testing"
)

func Test_canMerge(t *testing.T) {
	a1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}
	a2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}}
	a3 := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}}
	ar := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 4}},
	}

	b1 := &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 8}}
	b2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 6}}

	c1 := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}}
	c2 := &TreeNode{Val: 3}

	d1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}

	type args struct {
		trees []*TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{[]*TreeNode{a1, a2, a3}}, ar},
		{"2", args{[]*TreeNode{b1, b2}}, nil},
		{"3", args{[]*TreeNode{c1, c2}}, nil},
		{"4", args{[]*TreeNode{d1}}, d1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMerge(tt.args.trees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

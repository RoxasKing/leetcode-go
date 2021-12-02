package main

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	a1 := &TreeNode{Val: 4}
	a2 := &TreeNode{Val: 2}
	a3 := &TreeNode{Val: 7}
	a4 := &TreeNode{Val: 1}
	a5 := &TreeNode{Val: 3}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a2.Right = a5

	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{a1, 2}, a2},
		{"2", args{a1, 5}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

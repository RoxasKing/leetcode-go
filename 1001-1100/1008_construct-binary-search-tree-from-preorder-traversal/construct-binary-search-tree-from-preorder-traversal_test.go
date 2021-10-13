package main

import (
	"reflect"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"1",
			args{[]int{8, 5, 1, 7, 10, 12}},
			&TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 7}},
				Right: &TreeNode{Val: 10, Right: &TreeNode{Val: 12}},
			},
		},
		{"2", args{[]int{1, 3}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstFromPreorder(tt.args.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

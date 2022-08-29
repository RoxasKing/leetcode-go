package main

import (
	"reflect"
	"testing"
)

func Test_printTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"1", args{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}}, [][]string{{"", "1", ""}, {"2", "", ""}}},
		{
			"2",
			args{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}},
			[][]string{{"", "", "", "1", "", "", ""}, {"", "2", "", "", "", "3", ""}, {"", "", "4", "", "", "", ""}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

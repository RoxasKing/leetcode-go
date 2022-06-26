package main

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 3}
	t3 := &TreeNode{Val: 2}
	t4 := &TreeNode{Val: 5}
	t5 := &TreeNode{Val: 3}
	t6 := &TreeNode{Val: 9}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t3.Right = t6

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{t1}, []int{1, 3, 9}},
		{"2", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValues(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

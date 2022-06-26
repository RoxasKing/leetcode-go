package main

import (
	"reflect"
	"testing"
)

func Test_findFrequentTreeSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{&TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: -3}}}, []int{2, -3, 4}},
		{"2", args{&TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: -5}}}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFrequentTreeSum(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFrequentTreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

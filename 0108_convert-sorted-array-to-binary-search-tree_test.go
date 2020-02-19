package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"",
			args{[]int{-10, -3, 0, 5, 9}},
			&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

package codinginterviews

import (
	"reflect"
	"testing"
)

func Test_levelOrderIII(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				&TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			[][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderIII(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

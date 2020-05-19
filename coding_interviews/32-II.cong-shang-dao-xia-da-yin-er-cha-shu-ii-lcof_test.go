package codinginterviews

import (
	"reflect"
	"testing"
)

func Test_levelOrderII(t *testing.T) {
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
				{9, 20},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderII(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderII() = %v, want %v", got, tt.want)
			}
		})
	}
}

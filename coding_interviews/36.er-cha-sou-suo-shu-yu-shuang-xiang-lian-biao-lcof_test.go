package codinginterviews

import (
	"reflect"
	"testing"
)

func Test_treeToDoublyList(t *testing.T) {
	exp1before1 := &TreeNode{Val: 4}
	exp1before2 := &TreeNode{Val: 2}
	exp1before3 := &TreeNode{Val: 5}
	exp1before4 := &TreeNode{Val: 1}
	exp1before5 := &TreeNode{Val: 3}
	exp1before1.Left = exp1before2
	exp1before1.Right = exp1before3
	exp1before2.Left = exp1before4
	exp1before2.Right = exp1before5

	exp1after1 := &TreeNode{Val: 4}
	exp1after2 := &TreeNode{Val: 2}
	exp1after3 := &TreeNode{Val: 5}
	exp1after4 := &TreeNode{Val: 1}
	exp1after5 := &TreeNode{Val: 3}
	exp1after4.Left = exp1after3
	exp1after4.Right = exp1after2
	exp1after2.Left = exp1after4
	exp1after2.Right = exp1after5
	exp1after5.Left = exp1after2
	exp1after5.Right = exp1after1
	exp1after1.Left = exp1after5
	exp1after1.Right = exp1after3
	exp1after3.Left = exp1after1
	exp1after3.Right = exp1after4

	exp2before1 := &TreeNode{Val: 2}

	exp2after1 := &TreeNode{Val: 2}
	exp2after1.Left = exp2after1
	exp2after1.Right = exp2after1

	exp3before1 := &TreeNode{Val: 1}
	exp3before2 := &TreeNode{Val: 2}
	exp3before3 := &TreeNode{Val: 3}
	exp3before2.Left = exp3before1
	exp3before2.Right = exp3before3

	exp3after1 := &TreeNode{Val: 1}
	exp3after2 := &TreeNode{Val: 2}
	exp3after3 := &TreeNode{Val: 3}
	exp3after1.Left = exp3after3
	exp3after1.Right = exp3after2
	exp3after2.Left = exp3after1
	exp3after2.Right = exp3after3
	exp3after3.Left = exp3after2
	exp3after3.Right = exp3after1

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"1",
			args{exp1before1},
			exp1after4,
		},
		{
			"2",
			args{exp2before1},
			exp2after1,
		},
		{
			"3",
			args{nil},
			nil,
		},
		{
			"4",
			args{exp3before2},
			exp3after1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeToDoublyList(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeToDoublyList() = %v, want %v", got, tt.want)
			}
		})
	}
}

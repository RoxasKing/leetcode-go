package main

import (
	"strconv"
	"strings"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{3},
			[]string{
				"1,#,2,#,3,#,#",
				"1,#,3,2,#,#,#",
				"2,1,3,#,#,#,#",
				"3,1,#,#,2,#,#",
				"3,2,#,1,#,#,#",
			},
		},
	}

	codec := Constructor()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateTrees(tt.args.n)
			dict := make(map[string]bool)
			for _, tree := range got {
				dict[codec.serialize(tree)] = true
			}
			for _, tree := range tt.want {
				if !dict[tree] {
					t.Errorf("generateTrees() = %v, want %v", dict, tt.want)
				}
			}
		})
	}
}

// from 0297_serialize-and-deserialize-binary-tree
type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	q := []*TreeNode{root}
	var strs []string
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			strs = append(strs, "#")
		} else {
			strs = append(strs, strconv.Itoa(node.Val))
			q = append(q, node.Left, node.Right)
		}
	}
	return strings.Join(strs, ",")
}

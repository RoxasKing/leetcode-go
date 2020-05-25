package codinginterviews

import (
	"strconv"
	"strings"
)

/*
  请实现两个函数，分别用来序列化和反序列化二叉树。
*/

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "$,"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	bytes := strings.Split(data, ",")
	var index int
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if index == len(bytes)-1 {
			return nil
		}
		if bytes[index] == "$" {
			index++
			return nil
		}
		rootVal, _ := strconv.Atoi(bytes[index])
		index++
		return &TreeNode{
			Val:   rootVal,
			Left:  dfs(),
			Right: dfs(),
		}
	}
	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

package main

import (
	"strconv"
	"strings"
)

/*
  Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

  Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

  Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

  Example 1:
    Input: root = [1,2,3,null,null,4,5]
    Output: [1,2,3,null,null,4,5]

  Example 2:
    Input: root = []
    Output: []

  Example 3:
    Input: root = [1]
    Output: [1]

  Example 4:
    Input: root = [1,2]
    Output: [1,2]

  Constraints:
    1. The number of nodes in the tree is in the range [0, 10^4].
    2. -1000 <= Node.val <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS

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
			continue
		}
		strs = append(strs, strconv.Itoa(node.Val))
		q = append(q, node.Left, node.Right)
	}
	return strings.Join(strs, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	if len(arr) == 1 {
		return nil
	}
	rootVal, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: rootVal}
	arr = arr[1:]
	q := []*TreeNode{root}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		if arr[0] != "#" {
			val, _ := strconv.Atoi(arr[0])
			t.Left = &TreeNode{Val: val}
			q = append(q, t.Left)
		}
		if arr[1] != "#" {
			val, _ := strconv.Atoi(arr[1])
			t.Right = &TreeNode{Val: val}
			q = append(q, t.Right)
		}
		arr = arr[2:]
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

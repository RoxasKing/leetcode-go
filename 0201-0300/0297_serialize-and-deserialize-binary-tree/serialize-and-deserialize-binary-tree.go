package main

import (
	"strconv"
	"strings"
)

// Tags:
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

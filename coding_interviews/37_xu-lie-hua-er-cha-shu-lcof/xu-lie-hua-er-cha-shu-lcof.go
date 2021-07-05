package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	strs := []string{strconv.Itoa(root.Val)}
	q := []*TreeNode{root.Left, root.Right}
	for len(q) > 0 {
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
	if data == "#" {
		return nil
	}
	strs := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(strs[0])
	strs = strs[1:]
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	for len(strs) > 0 {
		node, left, right := queue[0], strs[0], strs[1]
		queue, strs = queue[1:], strs[2:]
		if left != "#" {
			lVal, _ := strconv.Atoi(left)
			node.Left = &TreeNode{Val: lVal}
			queue = append(queue, node.Left)
		}
		if right != "#" {
			rVal, _ := strconv.Atoi(right)
			node.Right = &TreeNode{Val: rVal}
			queue = append(queue, node.Right)
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

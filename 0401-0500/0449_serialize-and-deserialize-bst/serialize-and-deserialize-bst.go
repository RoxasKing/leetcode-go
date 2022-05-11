package main

import (
	"strconv"
	"strings"
)

// Difficulty:
// Medium

// Tags:
// BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	a := []string{}
	for q := []*TreeNode{root}; len(q) > 0; q = q[1:] {
		t := q[0]
		if t == nil {
			a = append(a, "#")
			continue
		}
		a = append(a, strconv.Itoa(t.Val))
		q = append(q, t.Left, t.Right)
	}
	return strings.Join(a, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	a := strings.Split(data, ",")
	if len(a) == 1 {
		return nil
	}
	v, _ := strconv.Atoi(a[0])
	o := &TreeNode{Val: v}
	for q := []*TreeNode{o}; len(q) > 0; q, a = q[1:], a[2:] {
		t := q[0]
		if a[1] != "#" {
			v, _ := strconv.Atoi(a[1])
			t.Left = &TreeNode{Val: v}
			q = append(q, t.Left)
		}
		if a[2] != "#" {
			v, _ := strconv.Atoi(a[2])
			t.Right = &TreeNode{Val: v}
			q = append(q, t.Right)
		}
	}
	return o
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

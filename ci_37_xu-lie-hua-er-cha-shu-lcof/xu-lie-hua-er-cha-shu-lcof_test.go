package main

import (
	"fmt"
)

func ExampleCodec() {
	c := Constructor()
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	str := c.serialize(root)
	fmt.Println(str)
	root = c.deserialize(str)
	str = c.serialize(root)
	fmt.Println(str)
	// Output:
	// 1,2,3,#,#,4,5,#,#,#,#
	// 1,2,3,#,#,4,5,#,#,#,#
}

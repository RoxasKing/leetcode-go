package main

import "fmt"

func ExampleCBTInserter() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	CBTI := Constructor(root)
	fmt.Println(CBTI.Insert(3))
	fmt.Println(CBTI.Insert(4))
	fmt.Println(CBTI.Get_root().Val)
	// Output:
	// 1
	// 2
	// 1
}

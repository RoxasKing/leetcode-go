package main

import "fmt"

func ExampleBSTIterator() {
	n0 := &TreeNode{Val: 7}
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 15}
	n0.Left, n0.Right = n1, n2
	n3 := &TreeNode{Val: 9}
	n4 := &TreeNode{Val: 20}
	n2.Left, n2.Right = n3, n4
	it := Constructor(n0)
	fmt.Println(it.Next())
	fmt.Println(it.Next())
	fmt.Println(it.HasNext())
	fmt.Println(it.Next())
	fmt.Println(it.HasNext())
	fmt.Println(it.Next())
	fmt.Println(it.HasNext())
	fmt.Println(it.Next())
	fmt.Println(it.HasNext())
	// Output:
	// 3
	// 7
	// true
	// 9
	// true
	// 15
	// true
	// 20
	// false
}

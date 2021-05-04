package main

import "fmt"

func ExampleFreqStack() {
	fs := Constructor()
	fs.Push(5)
	fs.Push(7)
	fs.Push(5)
	fs.Push(7)
	fs.Push(4)
	fs.Push(5)
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())

	fmt.Println()
	fs = Constructor()
	fs.Push(1)
	fs.Push(1)
	fs.Push(1)
	fs.Push(1)
	fs.Push(2)
	fs.Push(2)
	fs.Push(2)
	fs.Push(3)
	fs.Push(3)
	fmt.Println(fs.Pop()) // 1
	fmt.Println(fs.Pop()) // 2
	fmt.Println(fs.Pop()) // 1
	fmt.Println(fs.Pop()) // 3
	fmt.Println(fs.Pop()) // 2
	fmt.Println(fs.Pop()) // 1
	fmt.Println(fs.Pop()) // 3
	fmt.Println(fs.Pop()) // 2
	fmt.Println(fs.Pop()) // 1

	// Output:
	// 5
	// 7
	// 5
	// 4
	//
	// 1
	// 2
	// 1
	// 3
	// 2
	// 1
	// 3
	// 2
	// 1
}

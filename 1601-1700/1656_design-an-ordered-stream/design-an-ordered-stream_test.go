package main

import "fmt"

func ExampleOrderedStream() {
	os := Constructor(5)
	fmt.Println(os.Insert(3, "ccccc"))
	fmt.Println(os.Insert(1, "aaaaa"))
	fmt.Println(os.Insert(2, "bbbbb"))
	fmt.Println(os.Insert(5, "eeeee"))
	fmt.Println(os.Insert(4, "ddddd"))
	// Output:
	// []
	// [aaaaa]
	// [bbbbb ccccc]
	// []
	// [ddddd eeeee]
}

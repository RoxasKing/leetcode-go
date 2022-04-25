package main

import "fmt"

func ExampleCodec() {
	obj := Constructor()
	url := "https://leetcode.com/problems/design-tinyurl"
	fmt.Println(obj.decode(obj.encode(url)))
	// Output:
	// https://leetcode.com/problems/design-tinyurl
}

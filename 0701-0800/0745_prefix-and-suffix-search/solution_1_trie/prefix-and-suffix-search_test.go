package main

import "fmt"

func ExampleWordFilter() {
	wf := Constructor([]string{"apple"})
	fmt.Println(wf.F("a", "e"))
	fmt.Println(wf.F("b", ""))
	// Output:
	// 0
	// -1
}

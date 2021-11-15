package main

import "fmt"

func ExampleMapSum() {
	ms := Constructor()
	ms.Insert("apple", 3)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("ap", 3)
	fmt.Println(ms.Sum("ap"))
	// Output:
	// 3
	// 5
	// 8
}

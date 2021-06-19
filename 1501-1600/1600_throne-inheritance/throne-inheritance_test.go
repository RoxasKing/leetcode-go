package main

import "fmt"

func ExampleThroneInheritance() {
	t := Constructor("king")
	t.Birth("king", "andy")
	t.Birth("king", "bob")
	t.Birth("king", "catherine")
	t.Birth("andy", "matthew")
	t.Birth("bob", "alex")
	t.Birth("bob", "asha")
	fmt.Println(t.GetInheritanceOrder())
	t.Death("bob")
	fmt.Println(t.GetInheritanceOrder())
	// Output:
	// [king andy matthew bob alex asha catherine]
	// [king andy matthew alex asha catherine]
}

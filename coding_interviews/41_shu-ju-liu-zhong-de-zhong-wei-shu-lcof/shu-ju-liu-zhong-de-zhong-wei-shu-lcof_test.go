package main

import (
	"fmt"
)

func ExampleMedianFinder() {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	m.AddNum(3)
	m.AddNum(4)
	fmt.Println(m.FindMedian())
	m.AddNum(5)
	fmt.Println(m.FindMedian())
	m.AddNum(6)
	fmt.Println(m.FindMedian())
	// Output:
	// 2.5
	// 3
	// 3.5
}

package main

import "fmt"

func ExampleMedianFinder() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(3)
	fmt.Println(mf.FindMedian())
	// Output:
	// 1.5
	// 2
}

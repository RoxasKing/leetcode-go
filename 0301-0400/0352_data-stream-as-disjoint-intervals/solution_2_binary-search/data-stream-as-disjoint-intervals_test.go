package main

import "fmt"

func ExampleSummaryRanges() {
	sr := Constructor()
	sr.AddNum(1)
	fmt.Println(sr.GetIntervals())
	sr.AddNum(3)
	fmt.Println(sr.GetIntervals())
	sr.AddNum(7)
	fmt.Println(sr.GetIntervals())
	sr.AddNum(2)
	fmt.Println(sr.GetIntervals())
	sr.AddNum(6)
	fmt.Println(sr.GetIntervals())
	sr.AddNum(5)
	fmt.Println(sr.GetIntervals())
	sr.AddNum(4)
	fmt.Println(sr.GetIntervals())

	fmt.Println()
	sr2 := Constructor()
	sr2.AddNum(6)
	fmt.Println(sr2.GetIntervals())
	sr2.AddNum(0)
	fmt.Println(sr2.GetIntervals())

	// Output:
	// [[1 1]]
	// [[1 1] [3 3]]
	// [[1 1] [3 3] [7 7]]
	// [[1 3] [7 7]]
	// [[1 3] [6 7]]
	// [[1 3] [5 7]]
	// [[1 7]]
	//
	// [[6 6]]
	// [[0 0] [6 6]]
}

package main

import "fmt"

func ExampleSolution() {
	r, x0, y0 := 1.0, 0.0, 0.0
	es := Constructor(r, x0, y0)
	o := es.RandPoint()
	x, y := o[0], o[1]
	valid := func() bool {
		return (x-x0)*(x-x0)+(y-y0)*(y-y0) <= r*r
	}
	if valid() {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	o = es.RandPoint()
	x, y = o[0], o[1]
	if valid() {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	o = es.RandPoint()
	x, y = o[0], o[1]
	if valid() {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	// Output:
	// true
	// true
	// true
}

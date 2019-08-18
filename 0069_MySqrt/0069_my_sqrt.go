package main

import "fmt"

func mySqrt(x int) int {
	out := x
	for out*out > x {
		out = (out + x/out) / 2
	}
	return out
}

func main() {
	x := 4
	fmt.Println(mySqrt(x))
}

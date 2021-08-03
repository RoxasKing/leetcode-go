package main

// Tags:
// Math
// Binary Search

func minimumPerimeter(neededApples int64) int64 {
	x := int64(1)
	for ; 2*x*(x+1)*(2*x+1) < neededApples; x++ {
	}
	return x * 8
}

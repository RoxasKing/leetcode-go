package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	half := pow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

func main() {
	x := float64(2.00000)
	n := -2
	fmt.Println(myPow(x, n))
}

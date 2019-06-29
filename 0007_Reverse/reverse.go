package main

import "fmt"

func reverse(x int) int {
	var out int
	for i := 0; x/10 != 0; i++ {
		out = out*10 + x%10
		x /= 10
	}
	out = out*10 + x
	if out < -2147483648 || out > 2147483647 {
		return 0
	}
	return out
}

func main() {
	i := 12345
	fmt.Println(reverse(i))
}

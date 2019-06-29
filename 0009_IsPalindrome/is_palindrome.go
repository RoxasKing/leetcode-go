package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev, new := 0, x
	for new != 0 {
		rev = rev*10 + new%10
		new /= 10
	}
	if x == rev {
		return true
	}
	return false
}

func main() {
	i := 101
	fmt.Println(isPalindrome(i))
}

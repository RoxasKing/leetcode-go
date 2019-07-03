package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev, newx := 0, x
	for newx != 0 {
		rev = rev*10 + newx%10
		newx /= 10
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

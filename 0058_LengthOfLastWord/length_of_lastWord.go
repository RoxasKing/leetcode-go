package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else {
			if count == 0 {
				continue
			} else {
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if haystack == needle {
		return 0
	}
	size := len(needle)
	for i := 0; i < len(haystack); i++ {
		if i+size <= len(haystack) && haystack[i:i+size] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := ""
	needle := ""
	fmt.Println(strStr(haystack, needle))
}

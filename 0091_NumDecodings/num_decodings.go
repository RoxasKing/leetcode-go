package main

import "fmt"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	save := make([]int, len(s)+1)
	save[0], save[1] = 1, 1
	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			save[i] += save[i-1]
		}
		if s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6' {
			save[i] += save[i-2]
		}
	}
	return save[len(s)]
}

func main() {
	s := "226"
	fmt.Println(numDecodings(s))
}

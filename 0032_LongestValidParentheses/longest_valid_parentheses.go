package main

import "fmt"

func longestValidParentheses(s string) int {
	var slice []byte
	var position []int
	var out int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			slice = append(slice, s[i])
			position = append(position, i)
		} else if s[i] == ')' && len(slice) != 0 && slice[len(slice)-1] == '(' {
			slice = slice[:len(slice)-1]
			position = position[:len(position)-1]
		} else {
			//slice = append(slice, s[i])
			//position = append(position, i)
		}
	}
	if len(position) != 0 {
		if len(position) == 1 && (position[0] == 0 || position[0] == len(s)-1) {
			return len(s) - 1
		}
		pointer := -1
		for k, v := range position {
			size := 0
			if v >= 0 {
				if pointer != -1 {
					size = v - pointer - 1
				} else {
					size = v
				}
			}
			if size > out {
				out = size
			}
			pointer = v
			if k == len(position)-1 && out < len(s)-1-pointer {
				out = len(s) - 1 - pointer
			}
		}
		return out
	} else {
		return len(s)
	}
}

func main() {
	s := "()()("
	fmt.Println(longestValidParentheses(s))
}

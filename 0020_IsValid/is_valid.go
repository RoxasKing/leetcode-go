package main

import "fmt"

func isValid(s string) bool {
	var stack []rune
	for _, elem := range s {
		if elem == '(' || elem == '{' || elem == '[' {
			stack = append(stack, elem)
			continue
		}
		if elem == ')' {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
		if elem == ']' {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
		if elem == '}' {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	str := ""
	fmt.Println(isValid(str))
}

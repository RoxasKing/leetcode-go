package main

import "strconv"

// Stack
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		if !(t == "+" || t == "-" || t == "*" || t == "/") {
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
			continue
		}

		top := len(stack) - 1
		num1, num2 := stack[top-1], stack[top]
		stack = stack[:top-1]
		switch t {
		case "+":
			stack = append(stack, num1+num2)
		case "-":
			stack = append(stack, num1-num2)
		case "*":
			stack = append(stack, num1*num2)
		case "/":
			stack = append(stack, num1/num2)
		}
	}
	return stack[0]
}

package main

// Tags:
// Stack
func clumsy(N int) int {
	stack := []int{}
	op := '+'
	flg := 0
	for num := N; num > 0; num-- {
		switch op {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		case '*':
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, num1*num)
		case '/':
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, num1/num)
		}

		switch flg {
		case 0:
			op = '*'
		case 1:
			op = '/'
		case 2:
			op = '+'
		case 3:
			op = '-'
		}
		flg = (flg + 1) % 4
	}

	out := 0
	for _, num := range stack {
		out += num
	}
	return out
}

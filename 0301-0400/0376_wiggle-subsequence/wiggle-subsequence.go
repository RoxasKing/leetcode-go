package main

func wiggleMaxLength(nums []int) int {
	var stack []int
	var flg bool
	for _, num := range nums {
		if len(stack) > 0 && stack[len(stack)-1] == num ||
			len(stack) > 1 && (flg && stack[len(stack)-1] > num || !flg && stack[len(stack)-1] < num) {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
		if len(stack) > 1 {
			if stack[len(stack)-2] > stack[len(stack)-1] {
				flg = true
			} else {
				flg = false
			}
		}
	}
	return len(stack)
}

// Greedy
func wiggleMaxLength2(nums []int) int {
	out := 0
	a, b := -1<<31, -1<<31
	for _, num := range nums {
		if out > 0 && b == num || out > 1 && (a > b && b > num || a < b && b < num) {
			b = num
		} else {
			out++
			a, b = b, num
		}
	}
	return out
}

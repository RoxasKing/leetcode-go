package main

// Tags:
// Sliding Window
func maxSatisfied(customers []int, grumpy []int, X int) int {
	out, max := 0, 0
	for i, cur := 0, 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			out += customers[i]
		} else {
			cur += customers[i]
		}

		if i > X-1 && grumpy[i-X] == 1 {
			cur -= customers[i-X]
		}

		if i >= X-1 && cur > max {
			max = cur
		}
	}
	out += max
	return out
}

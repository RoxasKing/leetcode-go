package main

// Tags:
// Math

func lexicalOrder(n int) []int {
	out := make([]int, n)
	for i, num := 0, 1; i < n; i++ {
		out[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			if num >= n {
				num /= 10
			}
			for num++; num%10 == 0; num /= 10 {
			}
		}
	}
	return out
}

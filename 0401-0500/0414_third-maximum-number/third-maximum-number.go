package main

func thirdMax(nums []int) int {
	a, b, c := -1<<63, -1<<63, -1<<63
	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
		} else if num < a && num > b {
			b, c = num, b
		} else if num < b && num > c {
			c = num
		}
	}
	if c == -1<<63 {
		return a
	}
	return c
}

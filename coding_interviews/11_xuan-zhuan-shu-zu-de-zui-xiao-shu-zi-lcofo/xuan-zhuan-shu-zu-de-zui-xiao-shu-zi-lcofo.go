package main

// Tags:
// Binary Search
func minArray(numbers []int) int {
	for len(numbers) > 1 && numbers[0] == numbers[len(numbers)-1] {
		numbers = numbers[1:]
	}
	return numbers[findRotateIdx(numbers)]
}

func findRotateIdx(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		m := (l + r) / 2
		if numbers[m] > numbers[m+1] {
			return m + 1
		}
		if numbers[m] >= numbers[l] {
			l = m + 1
		} else {
			r = m
		}
	}
	return 0
}

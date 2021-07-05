package main

// Tags:
// Brain Teaser

func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	out := []int{0, 0}

	if a+2 == c {
		return out
	}

	if a+2 >= b || b+2 >= c {
		out[0] = 1
	} else {
		out[0] = 2
	}

	out[1] += c - a - 2

	return out
}

package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func totalFruit(fruits []int) int {
	o, l1, r1, l2, r2 := 0, -1, -1, -1, -1
	for i, x := range fruits {
		switch {
		case l1 == -1:
			l1, r1 = i, i
		case x == fruits[l1]:
			r1 = i
		case l2 == -1:
			l2, r2 = i, i
		case x == fruits[l2]:
			r2 = i
		default:
			o = max(o, max(r1, r2)+1-l1)
			l1, r1, l2, r2 = min(r1, r2)+1, max(r1, r2), i, i
		}
	}
	o = max(o, max(r1, r2)+1-l1)
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

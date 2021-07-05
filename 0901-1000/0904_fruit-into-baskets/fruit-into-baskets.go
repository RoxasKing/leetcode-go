package main

// Tags:
// Two Pointers
func totalFruit(tree []int) int {
	out := 0
	l1, l2, r1, r2 := 0, 0, 0, 0
	for i := range tree {
		if tree[i] == tree[l1] {
			l2 = i
		} else if tree[i] == tree[r1] {
			r2 = i
		} else if tree[l1] == tree[r1] {
			r1, r2 = i, i
		} else {
			out = Max(out, Max(l2, r2)+1-l1)
			l1, l2, r1, r2 = Min(l2+1, r2+1), Max(l2, r2), i, i
		}
	}
	out = Max(out, Max(l2, r2)+1-l1)
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

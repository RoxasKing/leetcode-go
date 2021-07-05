package main

func orchestraLayout(num int, xPos int, yPos int) int {
	i := Max(Max(num-1-xPos, xPos), Max(num-1-yPos, yPos))
	u, d, l, r := num-1-i, i, num-1-i, i
	idx := (num*num - (r+1-l)*(r+1-l)) % 9
	if xPos == u {
		idx = (idx + yPos + 1 - l) % 9
		if idx == 0 {
			return 9
		}
		return idx
	} else {
		idx = (idx + r + 1 - l) % 9
		u++
	}
	if yPos == r {
		idx = (idx + xPos + 1 - u) % 9
		if idx == 0 {
			return 9
		}
		return idx
	} else {
		idx = (idx + d + 1 - u) % 9
		r--
	}
	if xPos == d {
		idx = (idx + r + 1 - yPos) % 9
		if idx == 0 {
			return 9
		}
		return idx
	} else {
		idx = (idx + r + 1 - l) % 9
		d--
	}
	idx = (idx + d + 1 - xPos) % 9
	if idx == 0 {
		return 9
	}
	return idx
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

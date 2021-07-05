package main

// Tags:
// Two Pointers
func maxArea(height []int) int {
	out := 0
	for l, r := 0, len(height)-1; l < r; {
		out = Max(out, Min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return out
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

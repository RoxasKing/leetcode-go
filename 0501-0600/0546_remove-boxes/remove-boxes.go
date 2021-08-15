package main

// Tags:
// Dynamic Programming

func removeBoxes(boxes []int) int {
	f := [100][100][100]int{}
	var dp func(l, r, k int) int
	dp = func(l, r, k int) int {
		if l > r {
			return 0
		}
		if f[l][r][k] != 0 {
			return f[l][r][k]
		}
		f[l][r][k] = dp(l, r-1, 0) + (k+1)*(k+1)
		for i := l; i < r; i++ {
			if boxes[i] != boxes[r] {
				continue
			}
			f[l][r][k] = Max(f[l][r][k], dp(l, i, k+1)+dp(i+1, r-1, 0))
		}
		return f[l][r][k]
	}
	return dp(0, len(boxes)-1, 0)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

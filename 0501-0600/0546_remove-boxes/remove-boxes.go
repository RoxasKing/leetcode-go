package main

// Tags:
// Dynamic Programming
func removeBoxes(boxes []int) int {
	dp := [100][100][100]int{}
	var cal func(int, int, int) int
	cal = func(l, r, k int) int {
		if l > r {
			return 0
		}
		if dp[l][r][k] != 0 {
			return dp[l][r][k]
		}
		for l < r && boxes[r-1] == boxes[r] {
			r--
			k++
		}
		dp[l][r][k] = cal(l, r-1, 0) + (k+1)*(k+1)
		for i := l; i < r; i++ {
			if boxes[i] == boxes[r] {
				dp[l][r][k] = Max(dp[l][r][k], cal(l, i, k+1)+cal(i+1, r-1, 0))
			}
		}
		return dp[l][r][k]
	}
	return cal(0, len(boxes)-1, 0)
}

// Backtracking (time out)
func removeBoxes2(boxes []int) int {
	var max int
	var backtrack func([]int, int)
	backtrack = func(boxes []int, cur int) {
		if len(boxes) == 0 {
			max = Max(max, cur)
			return
		}
		var l, r int
		for l < len(boxes) {
			r = l
			for r+1 < len(boxes) && boxes[r+1] == boxes[r] {
				r++
			}
			newBoxes := make([]int, len(boxes)-(r+1-l))
			copy(newBoxes[:l], boxes[:l])
			copy(newBoxes[l:], boxes[r+1:])
			backtrack(newBoxes, cur+(r+1-l)*(r+1-l))
			l = r + 1
		}
	}
	backtrack(boxes, 0)
	return max
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func numsSameConsecDiff(n int, k int) []int {
	o := []int{}
	num := 0
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			o = append(o, num)
			return
		}
		for x := 0; x <= 9; x++ {
			if i == 0 && x == 0 || i > 0 && abs(num%10-x) != k {
				continue
			}
			num = num*10 + x
			backtrack(i + 1)
			num /= 10
		}
	}
	backtrack(0)
	return o
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

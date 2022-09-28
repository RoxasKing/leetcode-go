package main

import "strconv"

// Difficulty:
// Medium

// Tags:
// Math
// Backtracking

func rotatedDigits(n int) int {
	num := strconv.Itoa(n)
	arr := []int{0, 1, 2, 5, 6, 8, 9}
	o, tie, ok := 0, true, false
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(num) {
			if ok {
				o++
			}
			return
		}
		for _, x := range arr {
			if tie && x > int(num[i]-'0') {
				break
			}
			t1, t2 := tie, ok
			if x < int(num[i]-'0') {
				tie = false
			}
			if x == 2 || x == 5 || x == 6 || x == 9 {
				ok = true
			}
			backtrack(i + 1)
			tie, ok = t1, t2
		}
	}
	backtrack(0)
	return o
}

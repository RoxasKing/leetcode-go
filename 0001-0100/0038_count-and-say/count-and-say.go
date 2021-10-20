package main

import "strconv"

// Difficulty:
// Medium

func countAndSay(n int) string {
	x := "1"
	for k := 2; k <= n; k++ {
		t := ""
		p, c := x[0], 0
		for i := range x {
			if x[i] == p {
				c++
			} else {
				t += strconv.Itoa(c) + string(p)
				p, c = x[i], 1
			}
		}
		x = t + strconv.Itoa(c) + string(p)
	}
	return x
}

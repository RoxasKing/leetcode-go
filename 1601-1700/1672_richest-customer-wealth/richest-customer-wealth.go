package main

// Difficulty:
// Easy

func maximumWealth(accounts [][]int) int {
	o := 0
	for _, acc := range accounts {
		c := 0
		for _, cash := range acc {
			c += cash
		}
		if o < c {
			o = c
		}
	}
	return o
}

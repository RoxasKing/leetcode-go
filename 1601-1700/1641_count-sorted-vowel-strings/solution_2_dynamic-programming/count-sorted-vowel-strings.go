package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func countVowelStrings(n int) int {
	f := [5]int{1, 1, 1, 1, 1}
	for t := 2; t <= n; t++ {
		for i := 3; i >= 0; i-- {
			f[i] += f[i+1]
		}
	}
	o := 0
	for i := 0; i < 5; i++ {
		o += f[i]
	}
	return o
}

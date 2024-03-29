package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func countVowelStrings(n int) int {
	f := [5]int{1, 1, 1, 1, 1}
	for t := 2; t <= n+1; t++ {
		for i := 3; i >= 0; i-- {
			f[i] += f[i+1]
		}
	}
	return f[0]
}

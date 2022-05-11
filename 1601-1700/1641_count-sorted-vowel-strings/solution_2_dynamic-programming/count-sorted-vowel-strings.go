package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func countVowelStrings(n int) int {
	f0, f1 := [5]int{}, [5]int{}
	for i := 0; i < 5; i++ {
		f0[i] = i + 1
	}
	for k := 1; k < n; k++ {
		f1[0] = 1
		for i := 1; i < 5; i++ {
			f1[i] = f1[i-1] + f0[i]
		}
		f0, f1 = f1, f0
	}
	return f0[4]
}

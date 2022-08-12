package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func countVowelPermutation(n int) int {
	const mod int = 1e9 + 7
	g := [][]int{{1}, {0, 2}, {0, 1, 3, 4}, {2, 4}, {0}}
	f0, f1 := [5]int{1, 1, 1, 1, 1}, [5]int{}
	for ; n > 1; n-- {
		for i := 0; i < 5; i++ {
			for _, j := range g[i] {
				f1[i] = (f1[i] + f0[j]) % mod
			}
		}
		f0 = [5]int{}
		f0, f1 = f1, f0
	}
	o := 0
	for i := 0; i < 5; i++ {
		o = (o + f0[i]) % mod
	}
	return o
}

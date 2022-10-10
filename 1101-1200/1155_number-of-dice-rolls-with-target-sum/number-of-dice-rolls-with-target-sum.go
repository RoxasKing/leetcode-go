package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func numRollsToTarget(n int, k int, target int) int {
	f0, f1 := make([]int, target+1), make([]int, target+1)
	f0[0] = 1
	for ; n > 0; n-- {
		for i := 0; i <= target; i++ {
			f1[i] = 0
			for j := 1; j <= i && j <= k; j++ {
				f1[i] = (f1[i] + f0[i-j]) % (1e9 + 7)
			}
		}
		f0, f1 = f1, f0
	}
	return f0[target]
}

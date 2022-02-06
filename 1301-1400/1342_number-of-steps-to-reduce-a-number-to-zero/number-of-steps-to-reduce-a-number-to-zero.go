package main

// Difficulty:
// Easy

// Tags:
// Dynamic Programming

func numberOfSteps(num int) int {
	f := make([]int, num+1)
	for i := 1; i <= num; i++ {
		if i&1 == 1 {
			f[i] = f[i-1] + 1
		} else {
			f[i] = f[i>>1] + 1
		}
	}
	return f[num]
}

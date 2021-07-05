package main

// Tags:
// Dynamic Programming
func deleteAndEarn(nums []int) int {
	cnt := [1e4 + 1]int{}
	for _, num := range nums {
		cnt[num] += num
	}

	f := [1e4 + 1]int{}
	f[1] = cnt[1]
	for i := 2; i <= 1e4; i++ {
		f[i] = Max(f[i-1], f[i-2]+cnt[i])
	}
	return f[1e4]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

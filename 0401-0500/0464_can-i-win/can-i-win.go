package main

// Difficulty:
// Medium

// Tags:
// Memorization
// Dynamic Programming
// Bitmask
// Bit Manipulation

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	f := make([]int8, 1<<maxChoosableInteger)
	for i := range f {
		f[i] = -1
	}
	var dp func(int, int) int8
	dp = func(mask, sum int) int8 {
		if f[mask] != -1 {
			return f[mask]
		}
		f[mask] = 0
		for i := 0; i < maxChoosableInteger; i++ {
			if (mask>>i)&1 == 0 {
				continue
			}
			if sum+i+1 >= desiredTotal || dp(mask^(1<<i), sum+i+1) == 0 {
				f[mask] = 1
				break
			}
		}
		return f[mask]
	}
	return dp(1<<maxChoosableInteger-1, 0) == 1
}

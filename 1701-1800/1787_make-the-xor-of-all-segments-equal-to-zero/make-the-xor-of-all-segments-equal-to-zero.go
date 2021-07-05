package main

// Tags:
// Dynamic Programming
// Greedy algorithm
// Bit Manipulation

func minChanges(nums []int, k int) int {
	freq := make([][1 << 10]int, k)
	for i, num := range nums {
		freq[i%k][num]++
	}

	maxFreq, f0 := 0, freq[0]
	for j := 0; j < 1<<10; j++ {
		maxFreq = Max(maxFreq, freq[0][j])
	}
	keep, minFreq := maxFreq, maxFreq

	for i := 1; i < k; i++ {
		maxFreq, f1 := 0, [1 << 10]int{}
		for j := 0; j < 1<<10; j++ {
			if freq[i][j] == 0 {
				continue
			}
			maxFreq = Max(maxFreq, freq[i][j])
			for k := 0; k < 1<<10; k++ {
				f1[j^k] = Max(f1[j^k], f0[k]+freq[i][j])
			}
		}
		keep += maxFreq
		minFreq = Min(minFreq, maxFreq)
		f0 = f1
	}

	return len(nums) - Max(keep-minFreq, f0[0])
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

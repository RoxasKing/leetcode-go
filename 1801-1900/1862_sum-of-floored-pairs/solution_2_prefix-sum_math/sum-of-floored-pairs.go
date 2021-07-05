package main

// Tags:
// Prefix Sum + Math

func sumOfFlooredPairs(nums []int) int {
	upper := 0
	freq := map[int]int{}
	for _, num := range nums {
		upper = Max(upper, num)
		freq[num]++
	}

	pFreq := make([]int, upper+1)
	for i := 1; i <= upper; i++ {
		pFreq[i] = pFreq[i-1] + freq[i]
	}

	out := 0
	for y := 1; y <= upper; y++ {
		if freq[y] == 0 {
			continue
		}
		for d := 1; d*y <= upper; d++ {
			out += freq[y] * d * (pFreq[Min((d+1)*y-1, upper)] - pFreq[d*y-1])
		}
	}
	return out % (1e9 + 7)
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

package main

// Difficulty:
// Medium

// Tags:
// Math
// Counting

func threeSumMulti(arr []int, target int) int {
	freq := [101]int{}
	for _, x := range arr {
		freq[x]++
	}
	out := 0
	for i := 0; i*3 < target; i++ {
		if freq[i] == 0 {
			continue
		}
		for j := Max(i+1, target-100-i); i+j*2 < target; j++ {
			if freq[j] == 0 {
				continue
			}
			out += freq[i] * freq[j] * freq[target-i-j]
			out %= mod
		}
	}
	for i := 0; i*3 < target; i++ {
		if j := target - i*2; j < 101 && freq[i] > 1 {
			out += freq[i] * (freq[i] - 1) / 2 * freq[j]
			out %= mod
		}
		if (target-i)&1 == 1 {
			continue
		}
		if j := (target - i) / 2; j < 101 && freq[j] > 1 {
			out += freq[i] * freq[j] * (freq[j] - 1) / 2
			out %= mod
		}
	}
	if i := target / 3; target%3 == 0 && freq[i] > 2 {
		out += freq[i] * (freq[i] - 1) * (freq[i] - 2) / 6
		out %= mod
	}
	return out
}

const mod int = 1e9 + 7

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

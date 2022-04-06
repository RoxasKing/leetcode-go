package main

// Difficulty:
// Medium

// Tags:
// Math
// Counting

func threeSumMulti(arr []int, target int) int {
	mod := int(1e9 + 7)
	freq := [101]int{}
	for _, x := range arr {
		freq[x]++
	}
	out := 0
	for i := 0; i < target/3+1; i++ {
		if freq[i] == 0 {
			continue
		}
		for j := Max(i+1, target-100-i); i+j<<1 < target; j++ {
			if freq[j] == 0 {
				continue
			}
			out += freq[i] * freq[j] * freq[target-i-j]
			out %= mod
		}
	}
	for i := Max(0, target-200); i*3 < target; i++ {
		if freq[i] > 1 && target-i<<1 < 101 {
			out += freq[i] * (freq[i] - 1) / 2 * freq[target-i<<1]
			out %= mod
		}
		if (target-i)&1 == 1 {
			continue
		}
		if j := (target - i) >> 1; j < 101 && freq[j] > 1 {
			out += freq[i] * freq[j] * (freq[j] - 1) >> 1
			out %= mod
		}
	}
	if avg := target / 3; target%3 == 0 && freq[avg] > 2 {
		out += freq[avg] * (freq[avg] - 1) * (freq[avg] - 2) / 6
		out %= mod
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

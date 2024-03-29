package main

// Difficulty:
// Hard

// Tags:
// Rabin-Karp

func longestDupSubstring(S string) string {
	n := len(S)
	base, module := 26, int(1e9+7)
	nums := make([]int, n)
	for i := range S {
		nums[i] = int(S[i] - 'a')
	}
	var out string
	l, r := 1, n
	for l < r {
		L := l + (r-l)>>1
		if res := search(S, nums, n, base, L, module); res != "" {
			out = res
			l = L + 1
		} else {
			r = L
		}
	}
	return out
}

func search(S string, nums []int, n, base, L, modulus int) string {
	var hash int
	for i := 0; i < L; i++ {
		hash = (hash*base + nums[i]) % modulus
	}
	mark := map[int]string{hash: S[:L]}
	baseL := Pow(base, L-1, modulus)
	for i := 1; i < n+1-L; i++ {
		hash = (base*(hash-baseL*nums[i-1]%modulus+modulus)%modulus + nums[i-1+L]) % modulus
		if val, ok := mark[hash]; ok && val == S[i:i+L] {
			return S[i : i+L]
		}
		mark[hash] = S[i : i+L]
	}
	return ""
}

func Pow(base, n, module int) int {
	out := 1
	for n > 0 {
		if n&1 == 1 {
			out = (out * base) % module
		}
		base = (base * base) % module
		n >>= 1
	}
	return out
}

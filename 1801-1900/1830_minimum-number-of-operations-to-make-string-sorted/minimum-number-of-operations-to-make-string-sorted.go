package main

// Tags:
// Math
func makeStringSorted(s string) int {
	n := len(s)
	fac := make([]int, n+1)
	fac[0] = 1
	facinv := make([]int, n+1)
	facinv[0] = 1
	for i := 1; i < n; i++ {
		fac[i] = fac[i-1] * i % mod
		facinv[i] = quickMul(fac[i], mod-2) // 费马小定理
	}

	freq := [26]int{}
	for i := range s {
		freq[s[i]-'a']++
	}

	out := 0
	for i := 0; i < n-1; i++ {
		rank := 0
		idx := int(s[i] - 'a')
		for j := 0; j < idx; j++ {
			rank += freq[j]
		}
		cnt := rank * fac[n-1-i] % mod
		for j := 0; j < 26; j++ {
			cnt *= facinv[freq[j]]
			cnt %= mod
		}
		out += cnt
		out %= mod
		freq[idx]--
	}
	return out
}

func quickMul(x, n int) int {
	out := 1
	for n > 0 {
		if n&1 == 1 {
			out *= x
			out %= mod
		}
		x *= x
		x %= mod
		n >>= 1
	}
	return out
}

var mod = int(1e9 + 7)

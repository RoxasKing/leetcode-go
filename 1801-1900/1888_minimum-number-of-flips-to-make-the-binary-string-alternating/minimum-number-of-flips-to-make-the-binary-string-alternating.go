package main

// Tags:
// Prefix Sum
// Greedy

func minFlips(s string) int {
	n := len(s)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + (int(s[i]-'0') ^ (i & 1))
	}

	out := Min(pSum[n], n-pSum[n])
	for i := 1; i < n; i++ {
		cnt := 0
		if i&1 == 0 {
			cnt += pSum[n] - pSum[i]
		} else {
			cnt += (n - i) - (pSum[n] - pSum[i])
		}
		if (n-i)&1 == 0 {
			cnt += pSum[i]
		} else {
			cnt += i - pSum[i]
		}
		out = Min(out, Min(cnt, n-cnt))
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

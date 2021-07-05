package main

// Tags:
// Dynamic Programming
func findSubstringInWraproundString(p string) int {
	dp := [26]int{}
	k := 0
	for i := range p {
		if i > 0 && (int(p[i-1]-'a')+1)%26 == int(p[i]-'a') {
			k++
		} else {
			k = 1
		}
		dp[p[i]-'a'] = Max(dp[p[i]-'a'], k)
	}
	out := 0
	for i := 0; i < 26; i++ {
		out += dp[i]
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

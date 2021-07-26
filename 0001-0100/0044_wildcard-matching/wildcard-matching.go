package main

// Tags:
// Dynamic Programming
func isMatch(s string, p string) bool {
	ns, np := len(s), len(p)
	dp := make([][]bool, ns+1)
	for i := range dp {
		dp[i] = make([]bool, np+1)
	}
	dp[0][0] = true
	for i := 0; i < np && p[i] == '*'; i++ {
		dp[0][i+1] = true
	}
	for i := range s {
		for j := range p {
			if s[i] == p[j] || '?' == p[j] {
				dp[i+1][j+1] = dp[i][j]
			} else if '*' == p[j] {
				dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1]
			}
		}
	}
	return dp[ns][np]
}

// Greedy
func isMatch2(s string, p string) bool {
	for len(s) > 0 && len(p) > 0 && p[len(p)-1] != '*' {
		if s[len(s)-1] == p[len(p)-1] || '?' == p[len(p)-1] {
			s = s[:len(s)-1]
			p = p[:len(p)-1]
		} else {
			return false
		}
	}
	if p == "" {
		return s == ""
	}
	sIndex, pIndex := 0, 0
	sRecord, pRecord := -1, -1
	for sIndex < len(s) && pIndex < len(p) {
		if s[sIndex] == p[pIndex] || '?' == p[pIndex] {
			sIndex++
			pIndex++
		} else if '*' == p[pIndex] {
			pIndex++
			sRecord, pRecord = sIndex, pIndex
		} else if sRecord != -1 && sRecord+1 < len(s) {
			sRecord++
			sIndex, pIndex = sRecord, pRecord
		} else {
			return false
		}
	}
	for pIndex < len(p) {
		if p[pIndex] != '*' {
			return false
		}
		pIndex++
	}
	return true
}

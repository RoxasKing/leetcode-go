package main

// Difficulty:
// Medium

// Tags:
// Greedy
// Monotonic Stack

func removeDuplicateLetters(s string) string {
	freq := [26]int{}
	for i := range s {
		freq[s[i]-'a']++
	}
	vis := [26]bool{}
	stk := []byte{}
	top := func() int { return len(stk) - 1 }
	for i := range s {
		c := s[i]
		if vis[c-'a'] {
			freq[c-'a']--
			continue
		}
		for len(stk) > 0 && stk[top()] > c && freq[stk[top()]-'a'] > 1 {
			idx := stk[top()] - 'a'
			stk = stk[:top()]
			vis[idx] = false
			freq[idx]--
		}
		vis[c-'a'] = true
		stk = append(stk, c)
	}
	return string(stk)
}

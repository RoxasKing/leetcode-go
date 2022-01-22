package main

// Difficulty:
// Easy

// Tags:
// Two Pointers
// Hash

func removePalindromeSub(s string) int {
	out := 0
	for p := []byte(s); len(p) != 0; {
		out++
		t := []byte{}
		for l, r := 0, len(p)-1; l < r; l++ {
			if p[l] == p[r] {
				r--
			} else {
				t = append(t, p[l])
			}
		}
		p = t
	}
	has := [26]bool{}
	for i := range s {
		has[s[i]-'a'] = true
	}
	cnt := 0
	for i := 0; i < 26; i++ {
		if has[i] {
			cnt++
		}
	}
	if out > cnt {
		out = cnt
	}
	return out
}

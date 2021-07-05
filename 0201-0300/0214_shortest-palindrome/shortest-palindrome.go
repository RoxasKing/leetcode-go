package main

// Tags:
// KMP
func shortestPalindrome(s string) string {
	n := len(s)
	fail := make([]int, n)
	for i := 0; i < n; i++ {
		fail[i] = -1
	}
	for i := 1; i < n; i++ {
		j := fail[i-1]
		for j != -1 && s[j+1] != s[i] {
			j = fail[j]
		}
		if s[j+1] == s[i] {
			fail[i] = j + 1
		}
	}
	best := -1
	for i := n - 1; i >= 0; i-- {
		for best != -1 && s[best+1] != s[i] {
			best = fail[best]
		}
		if s[best+1] == s[i] {
			best++
		}
	}
	var prefix string
	if best != n-1 {
		prefix = s[best+1:]
	}
	return reverse(prefix) + s
}

func reverse(s string) string {
	letters := []byte(s)
	for i := 0; i < len(s)>>1; i++ {
		letters[i], letters[len(s)-1-i] = letters[len(s)-1-i], letters[i]
	}
	return string(letters)
}

// Hash
func shortestPalindrome2(s string) string {
	n := len(s)
	base, mod := 131, 1000000007
	l, r, mul := 0, 0, 1
	best := -1
	for i := 0; i < n; i++ {
		l = (l*base + int(s[i]-'0')) % mod
		r = (r + mul*int(s[i]-'0')) % mod
		if l == r {
			best = i
		}
		mul = mul * base % mod
	}
	var prefix string
	if best != n-1 {
		prefix = s[best+1:]
	}
	return reverse(prefix) + s
}

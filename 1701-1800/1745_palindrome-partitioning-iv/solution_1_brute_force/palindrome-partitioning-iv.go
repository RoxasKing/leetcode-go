package main

// Tags:
// Brute Force

func checkPartitioning(s string) bool {
	n := len(s)
	for i := 0; i < n-2; i++ {
		if !isPalindrome(s[:i+1]) {
			continue
		}
		for j := i + 1; j < n-1; j++ {
			if !isPalindrome(s[i+1 : j+1]) {
				continue
			}
			if isPalindrome(s[j+1:]) {
				return true
			}
		}
	}
	return false
}

func isPalindrome(s string) bool {
	n := len(s)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

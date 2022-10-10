package main

// Difficulty:
// Medium

// Tags:
// Greedy

func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	a := []byte(palindrome)
	for i := range a {
		if a[i] > 'a' && (n&1 == 0 || i != n/2) {
			a[i] = 'a'
			return string(a)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if a[i] < 'z' && (n&1 == 0 || i != n/2) {
			a[i]++
			return string(a)
		}
	}
	return ""
}

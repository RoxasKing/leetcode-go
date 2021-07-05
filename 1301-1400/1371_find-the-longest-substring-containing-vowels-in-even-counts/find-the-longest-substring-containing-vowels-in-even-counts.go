package main

// Tags:
// Prefix Sum + Bit Manipulation
func findTheLongestSubstring(s string) int {
	pos := make([]int, 1<<5)
	pos[0] = -1
	for i := 1; i < 1<<5; i++ {
		pos[i] = -2
	}
	out, xor := 0, 0
	for i := range s {
		switch s[i] {
		case 'a':
			xor ^= 1 << 0
		case 'e':
			xor ^= 1 << 1
		case 'i':
			xor ^= 1 << 2
		case 'o':
			xor ^= 1 << 3
		case 'u':
			xor ^= 1 << 4
		}
		if j := pos[xor]; j == -2 {
			pos[xor] = i
		} else if i-j > out {
			out = i - j
		}
	}
	return out
}

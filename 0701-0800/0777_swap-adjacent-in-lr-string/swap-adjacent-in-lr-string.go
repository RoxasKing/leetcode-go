package main

// Difficulty:
// Medium

// Tags:
// Brain Teaser

func canTransform(start string, end string) bool {
	for n, i, j := len(start), -1, -1; ; {
		for i++; i < n && start[i] == 'X'; i++ {
		}
		for j++; j < n && end[j] == 'X'; j++ {
		}
		if i == n && j == n {
			break
		}
		if i == n || j == n || start[i] != end[j] || end[j] == 'L' && i < j || end[j] == 'R' && i > j {
			return false
		}
	}
	return true
}

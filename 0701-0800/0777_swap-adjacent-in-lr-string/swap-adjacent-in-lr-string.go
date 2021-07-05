package main

// Tags:
// Brain Teaser

func canTransform(start string, end string) bool {
	for i, j, n := -1, -1, len(start); ; {
		for i++; i < n && start[i] == 'X'; i++ {
		}
		for j++; j < n && end[j] == 'X'; j++ {
		}

		if i == n && j == n {
			break
		}

		if i == n || j == n || start[i] != end[j] || start[i] == 'L' && i < j || start[i] == 'R' && i > j {
			return false
		}
	}
	return true
}

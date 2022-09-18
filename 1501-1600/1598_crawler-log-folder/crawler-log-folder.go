package main

// Difficulty:
// Easy

func minOperations(logs []string) int {
	cnt := 0
	for _, s := range logs {
		if s == "./" {
			continue
		}
		if s == "../" {
			cnt = max(0, cnt-1)
			continue
		}
		cnt++
	}
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

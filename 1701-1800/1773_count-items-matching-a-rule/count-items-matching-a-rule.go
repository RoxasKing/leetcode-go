package main

// Difficulty:
// Easy

// Tags:
// Hash

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	o, h := 0, map[string]int{"type": 0, "color": 1, "name": 2}
	for _, item := range items {
		if item[h[ruleKey]] == ruleValue {
			o++
		}
	}
	return o
}

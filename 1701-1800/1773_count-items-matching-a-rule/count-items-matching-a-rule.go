package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	out := 0
	for _, item := range items {
		if ruleKey == "type" && item[0] == ruleValue ||
			ruleKey == "color" && item[1] == ruleValue ||
			ruleKey == "name" && item[2] == ruleValue {
			out++
		}
	}
	return out
}

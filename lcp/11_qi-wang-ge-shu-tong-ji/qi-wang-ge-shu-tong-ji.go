package main

func expectNumber(scores []int) int {
	out := 0
	visited := make(map[int]bool)
	for _, score := range scores {
		if visited[score] {
			continue
		}
		visited[score] = true
		out++
	}
	return out
}

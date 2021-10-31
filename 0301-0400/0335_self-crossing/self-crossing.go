package main

// Difficulty:
// Hard

// Tags:
// Math

func isSelfCrossing(dist []int) bool {
	n := len(dist)
	if n < 4 {
		return false
	}
	for i := 3; i < n; i++ {
		if i >= 3 && dist[i-1] <= dist[i-3] && dist[i] >= dist[i-2] ||
			i >= 4 && dist[i-3] == dist[i-1] && dist[i]+dist[i-4] >= dist[i-2] ||
			i >= 5 && dist[i-2] > dist[i-4] && dist[i-3] > dist[i-1] &&
				dist[i]+dist[i-4] >= dist[i-2] && dist[i-1]+dist[i-5] >= dist[i-3] {
			return true
		}
	}
	return false
}
